package rules

import (
	"github.com/yoheimuta/go-protoparser/parser"
	"github.com/yoheimuta/protolint/internal/addon/rules/internal/visitor"
	"github.com/yoheimuta/protolint/internal/linter/report"
	"github.com/yoheimuta/protolint/internal/strs"
)

// FieldNamesLowerSnakeCaseRule verifies that all field names are underscore_separated_names.
// See https://developers.google.com/protocol-buffers/docs/style#message-and-field-names.
type FieldNamesLowerSnakeCaseRule struct{}

// NewFieldNamesLowerSnakeCaseRule creates a new FieldNamesLowerSnakeCaseRule.
func NewFieldNamesLowerSnakeCaseRule() FieldNamesLowerSnakeCaseRule {
	return FieldNamesLowerSnakeCaseRule{}
}

// ID returns the ID of this rule.
func (r FieldNamesLowerSnakeCaseRule) ID() string {
	return "FIELD_NAMES_LOWER_SNAKE_CASE"
}

// Purpose returns the purpose of this rule.
func (r FieldNamesLowerSnakeCaseRule) Purpose() string {
	return "Verifies that all field names are underscore_separated_names."
}

// IsOfficial decides whether or not this rule belongs to the official guide.
func (r FieldNamesLowerSnakeCaseRule) IsOfficial() bool {
	return true
}

// Apply applies the rule to the proto.
func (r FieldNamesLowerSnakeCaseRule) Apply(proto *parser.Proto) ([]report.Failure, error) {
	v := &fieldNamesLowerSnakeCaseVisitor{
		BaseAddVisitor: visitor.NewBaseAddVisitor(),
	}
	return visitor.RunVisitor(v, proto, r.ID())
}

type fieldNamesLowerSnakeCaseVisitor struct {
	*visitor.BaseAddVisitor
}

// VisitField checks the field.
func (v *fieldNamesLowerSnakeCaseVisitor) VisitField(field *parser.Field) bool {
	if !strs.IsLowerSnakeCase(field.FieldName) {
		v.AddFailuref(field.Meta.Pos, "Field name %q must be LowerSnakeCase", field.FieldName)
	}
	return false
}
