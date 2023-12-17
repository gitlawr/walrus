package cmd

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"os/exec"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/seal-io/walrus/pkg/cli/api"
	"github.com/seal-io/walrus/pkg/cli/ask"
	"github.com/seal-io/walrus/pkg/cli/common"
	"github.com/seal-io/walrus/pkg/cli/config"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/strs"
)

// NewConfigCmd generate config command.
func NewConfigCmd(serverConfig *config.Config, root *cobra.Command) *cobra.Command {
	// Command config setup.
	cfg := config.ServerContext{}

	// Command config setup.
	setupCmd := &cobra.Command{
		Use:   "setup short-name",
		Short: "Connect Walrus server and setup cli",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Configuration value from environment variables.
			viper.SetEnvPrefix("WALRUS")
			viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			viper.AutomaticEnv()
			common.BindFlags(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			isInputByFlags := common.InputByFlags(cmd)

			// When the user does not provide any flags, interactive configuration is provided.
			if !isInputByFlags {
				qs := questions(serverConfig)
				err := survey.Ask(
					qs,
					&cfg,
					survey.WithHideCharacter('*'),
				)
				if err != nil {
					panic(err)
				}
			}

			set := cmd.Flags()
			err := setup(cfg, serverConfig, set, isInputByFlags)
			if err != nil {
				panic(err)
			}
		},
	}

	cfg.AddFlags(setupCmd)

	// Command config sync.
	syncCmd := &cobra.Command{
		Use:   "sync short-name",
		Short: "Sync cli action to the latest",
		Run: func(cmd *cobra.Command, args []string) {
			err := sync(serverConfig, root)
			if err != nil {
				panic(err)
			}
		},
	}

	// Command config current context.
	currentContextCmd := &cobra.Command{
		Use:   "current-context short-name",
		Short: "Get current context",
		Run: func(cmd *cobra.Command, args []string) {
			currentContext(serverConfig)
		},
	}

	// Command config.
	configCmd := &cobra.Command{
		Use:     "config",
		Short:   "Manage CLI configuration",
		GroupID: common.GroupOther.ID,
	}
	configCmd.AddCommand(
		setupCmd,
		syncCmd,
		currentContextCmd,
	)

	return configCmd
}

// currentContext define the function for config current-context command.
func currentContext(serverConfig *config.Config) {
	if serverConfig.Project != "" {
		name := serverConfig.Project
		if name != "" {
			fmt.Println("Current Project: " + name)
		}

		env := serverConfig.Environment
		if env != "" {
			fmt.Println("Current Environment: " + env)
		}
	}
}

// setup define the function for config setup command.
func setup(sc config.ServerContext, serverConfig *config.Config, flags *pflag.FlagSet, inputByFlags bool) error {
	merged := config.Config{
		ServerContext: sc,
	}

	if inputByFlags {
		origin := &config.ServerContext{}
		if serverConfig.Server != "" {
			origin = &serverConfig.ServerContext
		}

		merged = config.Config{
			ServerContext: origin.Merge(sc, flags),
		}
	} else {
		// Ignore empty.
		if merged.Project == `""` {
			merged.Project = ""
		}

		if merged.Environment == `""` {
			merged.Environment = ""
		}
	}
	isLocal, err := flags.GetBool("local")
	if err != nil {
		return err
	}
	if isLocal {
		installed, err := checkWalrusDockerExtension()
		if err != nil {
			return err
		}
		if !installed {
			confirm := ""
			prompt := &survey.Input{
				Message: "Install Walrus docker extension to proceed [y/N]",
			}

			survey.AskOne(prompt, &confirm)
			if confirm != "y" {
				return nil
			}

			err = installWalrusDockerExtension()
			if err != nil {
				return fmt.Errorf("failed to install walrus docker extension: %w", err)
			}
			fmt.Println("Walrus docker extension installed successfully.")
			fmt.Println("Checking readiness...")
		}
		merged.ServerContext.Server = "https://localhost:7443"
		merged.ServerContext.Insecure = true
		merged.ServerContext.Project = "default"
		merged.ServerContext.Environment = "local"
	}

	wait.PollUntilContextTimeout(context.Background(), 5*time.Second, 5*time.Minute, false, func(ctx context.Context) (done bool, err error) {
		err = merged.ValidateAndSetup()
		if err != nil {
			return false, nil
		}
		return true, nil
	})

	serverConfig.ServerContext = merged.ServerContext

	err = api.InitOpenAPI(serverConfig, true)
	if err != nil {
		return err
	}

	fmt.Println("Walrus CLI is configured.")

	return config.SetServerContextToCache(serverConfig.ServerContext)
}

// sync define the function for config sync command.
func sync(serverConfig *config.Config, root *cobra.Command) error {
	err := serverConfig.ValidateAndSetup()
	if err != nil {
		return err
	}

	err = Load(serverConfig, root, true)

	return err
}

// Load OpenAPI from cache or remote and setup command.
func Load(sc *config.Config, root *cobra.Command, skipCache bool) error {
	start := time.Now()
	defer func() {
		log.Debugf("API loading took %s", time.Since(start))
	}()

	err := api.InitOpenAPI(sc, skipCache)
	if err != nil {
		return err
	}

	api.OpenAPI.GenerateCommand(sc, root)

	return err
}

// questions is the interactive prompt questions use to config CLI.
func questions(serverConfig *config.Config) []*survey.Question {
	hiddenPassword := func(val string) string {
		return fmt.Sprintf("****%s", strs.LastContent(val, 4))
	}

	ap := &ask.Password{
		Message:        strs.Question(config.FlagNameToken),
		Default:        serverConfig.Token,
		DefaultDisplay: hiddenPassword(serverConfig.Token),
	}

	proj := serverConfig.Project
	if proj == "" {
		proj = "default"
	}

	qs := []*survey.Question{
		{
			Name: config.FlagNameServer,
			Prompt: &survey.Input{
				Message: strs.Question(config.FlagNameServer),
				Default: serverConfig.Server,
			},
			Validate: survey.Required,
		},
		{
			Name:     config.FlagNameToken,
			Prompt:   ap,
			Validate: ap.Required,
		},
		{
			Name: config.FlagNameInsecure,
			Prompt: &survey.Confirm{
				Message: strs.Question(config.FlagNameInsecure),
				Default: serverConfig.Insecure,
			},
		},
		{
			Name: config.FlagNameProject,
			Prompt: &survey.Input{
				Message: strs.Question(config.FlagNameProject),
				Default: proj,
			},
		},
		{
			Name: config.FlagNameEnvironment,
			Prompt: &survey.Input{
				Message: strs.Question(config.FlagNameEnvironment),
				Default: serverConfig.Environment,
			},
		},
	}

	return qs
}

func checkWalrusDockerExtension() (bool, error) {
	cmd := exec.Command("docker", "extension", "ls")
	out, err := cmd.Output()
	if strings.Contains(string(out), "walrus-docker-extension") {
		return true, nil
	}
	return false, err
}

func installWalrusDockerExtension() error {
	cmd := exec.Command("docker", "extension", "install", "lawr/walrus-docker-extension:latest", "--force")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
