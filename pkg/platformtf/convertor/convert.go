package convertor

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/platformtf/block"
	"github.com/seal-io/seal/utils/log"
)

const (
	// ProviderCustom is the custom(not actually exist) provider. which is used to
	// generate the provider block for the custom connector.
	ProviderCustom = "custom"
	ProviderK8s    = "kubernetes"
	ProviderHelm   = "helm"
	// Add more providers
)

type ConvertOptions struct {
	SecretMountPath string
	ConnSeparator   string
	Providers       []string
}

// _providersToValidate if the required providers contains any
// of the _providersToValidates, it must be contained of the
// generated Blocks labels.
var _providersToValidate []string

// _providerConvertors mutate the connector to provider block.
var _providerConvertors = make(map[string]Convertor, 0)

func init() {
	convertors := []Convertor{
		CustomConvertor{},
		K8sConvertor{},
		HelmConvertor{},
		// add more convertors
	}
	for _, c := range convertors {
		_providerConvertors[c.ProviderType()] = c
		// custom provider is not a real provider in terraform,
		// validating custom provider is meaningless.
		if c.ProviderType() == ProviderCustom {
			continue
		}
		_providersToValidate = append(_providersToValidate, c.ProviderType())
	}
}

// LoadConvertor loads the convertor by the provider type.
func LoadConvertor(provider string) Convertor {
	if _, ok := _providerConvertors[provider]; !ok {
		return nil
	}

	return _providerConvertors[provider]
}

// ToProvidersBlocks converts the connectors to provider blocks with required providers.
func ToProvidersBlocks(providers []string, connectors model.Connectors, opts ConvertOptions) (blocks block.Blocks, err error) {
	// customProviders is the providers generated from the custom connectors.
	var customProviders []string
	for _, p := range providers {
		var convertBlocks block.Blocks
		convertBlocks, err = ToProviderBlocks(p, connectors, opts)
		if err != nil {
			return nil, err
		}
		if convertBlocks == nil {
			continue
		}
		blocks = append(blocks, convertBlocks...)
		// record the terraform provider name that generated from the custom connectors.
		// e.g. connectors have types helm, aws, alicloud, etc.
		// we need to validate with the required providers.
		if p == ProviderCustom {
			customProviders, err = convertBlocks.GetProviderNames()
			if err != nil {
				return nil, err
			}
		}
	}

	// supportedProviders is the providers which can be generated by the convertors.
	supportedProviders := append([]string{}, _providersToValidate...)
	supportedProviders = append(supportedProviders, customProviders...)
	ok, err := validateRequiredProviders(providers, supportedProviders, blocks)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("failed to validate providers: %v, current blockProviders: %v", providers, blocks)
	}

	return blocks, nil
}

// ToProviderBlocks converts the connectors to blocks with provider name.
func ToProviderBlocks(provider string, connectors model.Connectors, opts ConvertOptions) (block.Blocks, error) {
	var (
		toBlockOpts Options
		conns       model.Connectors
	)
	switch provider {
	case ProviderK8s, ProviderHelm:
		toBlockOpts = K8sConvertorOptions{
			ConnSeparator: opts.ConnSeparator,
			ConfigPath:    opts.SecretMountPath,
		}
	case ProviderCustom:
		toBlockOpts = CustomConvertorOptions{
			Providers: opts.Providers,
		}
	default:
		// TODO add more options
	}

	convertor := LoadConvertor(provider)
	if convertor == nil {
		return nil, nil
	}
	conns = convertor.GetConnectors(connectors)

	return convertor.ToBlocks(conns, toBlockOpts)
}

// validateRequiredProviders the providers both in providers and supportedProviders
// need to be checked in the generated blocks.
func validateRequiredProviders(providers, supportedProviders []string, blocks block.Blocks) (bool, error) {
	logger := log.WithName("platformtf").WithName("convert")
	// blockProviders is the providers of the generated blocks.
	blockProviders, err := blocks.GetProviderNames()
	if err != nil {
		return false, err
	}
	actual := sets.NewString(blockProviders...)

	// supported is the providers which convertors can generate with the given connectors.
	// custom providers are the providers which are generated by the primary provider.
	supported := sets.NewString(supportedProviders...)

	// get the intersection of the required and supported providers to validate.
	required := sets.NewString(providers...)
	expected := required.Intersection(supported)

	if actual.IsSuperset(expected) {
		// remove the blocks which are not in the required providers.
		for i := 0; i < len(blocks); i++ {
			if len(blocks[i].Labels) == 0 {
				logger.Errorf("invalid block: %v", blocks[i])
				return false, nil
			}
			providerType := blocks[i].Labels[0]
			if !required.Has(providerType) {
				blocks.Remove(blocks[i])
				i--
			}
		}
		return true, nil
	}

	return false, nil
}