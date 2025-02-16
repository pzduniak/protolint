package config

// Lint represents the lint configuration.
type Lint struct {
	Ignores     Ignores
	Directories Directories
	Rules       Rules
	RulesOption RulesOption `yaml:"rules_option"`
}

// ExternalConfig represents the external configuration.
type ExternalConfig struct {
	Lint Lint
}

// ShouldSkipRule checks whether to skip applying the rule to the file.
func (c ExternalConfig) ShouldSkipRule(
	ruleID string,
	displayPath string,
	defaultRuleIDs []string,
) bool {
	lint := c.Lint
	return lint.Ignores.shouldSkipRule(ruleID, displayPath) ||
		lint.Directories.shouldSkipRule(displayPath) ||
		lint.Rules.shouldSkipRule(ruleID, defaultRuleIDs)
}
