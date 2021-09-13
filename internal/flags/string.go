package flags

type StringFlag struct {
	Value         string
	originalValue string
	isSet         bool
}

func (f *StringFlag) SetValue(s string) error {
	f.originalValue = s
	// i'm not aware of any validation around names needed for string flags like the name
	f.Value = s
	f.isSet = true
	return nil
}

func (f *StringFlag) OriginalValue() string {
	return f.originalValue
}

func (f *StringFlag) IsSet() bool {
	return f.isSet
}
func (f *StringFlag) Type() string {
	return "string"
}
