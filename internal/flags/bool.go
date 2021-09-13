package flags

import (
	"errors"
)

type BoolFlag struct {
	Value         bool
	originalValue string
	isSet         bool
}

func (f *BoolFlag) SetValue(s string) error {
	f.originalValue = s
	f.isSet = true
	if s != "" {
		return errors.New("boolean flag cannot have a value")
	}
	return nil
}

func (f *BoolFlag) OriginalValue() string {
	return f.originalValue
}

func (f *BoolFlag) IsSet() bool {
	return f.isSet
}

func (f *BoolFlag) Type() string {
	return "bool"
}
