package types

type Selector struct {
	EnvironmentNames  []string          `json:"environmentNames,omitempty"`
	EnvironmentTypes  []string          `json:"environmentTypes,omitempty"`
	ResourceLabels    map[string]string `json:"resourceLabels,omitempty"`
	EnvironmentLabels map[string]string `json:"environmentLabels,omitempty"`
}
