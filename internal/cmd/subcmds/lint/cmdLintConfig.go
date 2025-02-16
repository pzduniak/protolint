package lint

import (
	"github.com/yoheimuta/protolint/internal/cmd/subcmds"
	"github.com/yoheimuta/protolint/internal/linter/config"
	"github.com/yoheimuta/protolint/internal/linter/file"
	"github.com/yoheimuta/protolint/internal/linter/rule"
)

// CmdLintConfig is a config for lint command.
type CmdLintConfig struct {
	external config.ExternalConfig
	fixMode  bool
	verbose  bool
}

// NewCmdLintConfig creates a new CmdLintConfig.
func NewCmdLintConfig(
	externalConfig config.ExternalConfig,
	flags Flags,
) CmdLintConfig {
	return CmdLintConfig{
		external: externalConfig,
		fixMode:  flags.FixMode,
		verbose:  flags.Verbose,
	}
}

// GenRules generates rules which are applied to the filename path.
func (c CmdLintConfig) GenRules(
	f file.ProtoFile,
) ([]rule.HasApply, error) {
	allRules := subcmds.NewAllRules(c.external.Lint.RulesOption, c.fixMode)

	var defaultRuleIDs []string
	if c.external.Lint.Rules.AllDefault {
		defaultRuleIDs = allRules.IDs()
	} else {
		defaultRuleIDs = allRules.Default().IDs()
	}

	var hasApplies []rule.HasApply
	for _, r := range allRules {
		if c.external.ShouldSkipRule(r.ID(), f.DisplayPath(), defaultRuleIDs) {
			continue
		}
		hasApplies = append(hasApplies, r)
	}

	return hasApplies, nil
}
