package resourcedefinitions

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/seal-io/walrus/pkg/dao/model"
)

func Match(
	matchingRules []*model.ResourceDefinitionMatchingRule,
	environmentName, environmentType string,
	environmentLabels, resourceLabels map[string]string,
) *model.ResourceDefinitionMatchingRule {
	for _, rule := range matchingRules {
		switch {
		case len(rule.Selector.EnvironmentNames) > 0 &&
			!slices.Contains(rule.Selector.EnvironmentNames, environmentName):
			continue
		case len(rule.Selector.EnvironmentTypes) > 0 &&
			!slices.Contains(rule.Selector.EnvironmentTypes, environmentType):
			continue
		case !matchLabels(rule.Selector.EnvironmentLabels, environmentLabels):
			continue
		case !matchLabels(rule.Selector.ResourceLabels, resourceLabels):
			continue
		default:
			return rule
		}
	}

	return nil
}

// MatchEnvironment returns the matching rule that pairs with the environment regardless of resource labels.
func MatchEnvironment(
	matchingRules []*model.ResourceDefinitionMatchingRule,
	environmentName, environmentType string,
	environmentLabels map[string]string,
) *model.ResourceDefinitionMatchingRule {
	for _, rule := range matchingRules {
		switch {
		case len(rule.Selector.EnvironmentNames) > 0 &&
			!slices.Contains(rule.Selector.EnvironmentNames, environmentName):
			continue
		case len(rule.Selector.EnvironmentTypes) > 0 &&
			!slices.Contains(rule.Selector.EnvironmentTypes, environmentType):
			continue
		case !matchLabels(rule.Selector.EnvironmentLabels, environmentLabels):
			continue
		default:
			return rule
		}
	}

	return nil
}

func matchLabels(selectors, labels map[string]string) bool {
	if len(selectors) == 0 {
		return true
	}

	for key, value := range selectors {
		if labelValue, exists := labels[key]; !exists || labelValue != value {
			return false
		}
	}

	return true
}

// IsConflict returns an error if two sets of project names conflict.
func IsConflict(pNames1, pNames2 []string) error {
	if len(pNames1) == 0 && len(pNames2) == 0 {
		return errors.New("both are applicable to all projects")
	}

	for _, pName1 := range pNames1 {
		for _, pName2 := range pNames2 {
			if pName1 == pName2 {
				return fmt.Errorf("both are applicable to project %q", pName1)
			}
		}
	}

	return nil
}
