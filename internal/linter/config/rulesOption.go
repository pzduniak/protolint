package config

// RulesOption represents the option for some rules.
type RulesOption struct {
	FileNamesLowerSnakeCase         FileNamesLowerSnakeCaseOption         `yaml:"file_names_lower_snake_case"`
	ImportsSorted                   ImportsSortedOption                   `yaml:"imports_sorted"`
	MaxLineLength                   MaxLineLengthOption                   `yaml:"max_line_length"`
	Indent                          IndentOption                          `yaml:"indent"`
	EnumFieldNamesZeroValueEndWith  EnumFieldNamesZeroValueEndWithOption  `yaml:"enum_field_names_zero_value_end_with"`
	ServiceNamesEndWith             ServiceNamesEndWithOption             `yaml:"service_names_end_with"`
	FieldNamesExcludePrepositions   FieldNamesExcludePrepositionsOption   `yaml:"field_names_exclude_prepositions"`
	MessageNamesExcludePrepositions MessageNamesExcludePrepositionsOption `yaml:"message_names_exclude_prepositions"`
}
