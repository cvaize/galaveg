package field

type View struct {
	FieldClass string
	Errors     []string
	Style      string

	Label      string
	LabelClass string

	Type        string
	Placeholder string

	InputClass string
	InputStyle string

	Form  string
	Name  string
	Value string

	Multiple  bool
	Required  bool
	Readonly  bool
	Autofocus bool

	Autocomplete   string
	Autocorrect    string
	Autocapitalize string

	Min    string
	Max    string
	Step   string
	Accept string

	// Для select: список простых значений
	OptionsValues []string

	// Для select: список структур/мап
	Options        []map[string]string
	OptionValueKey string
	OptionLabelKey string
}
