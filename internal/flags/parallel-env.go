package flags

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type PEName int64

const (
	undefined PEName = iota
	orte
	smp
)

func (p PEName) String() string {
	switch p {
	// i'm somewhat not sure how best to differentiate
	// undefined vs unknown, especially since
	// it should likely never happen and instead error
	case undefined:
		return "undefined"
	case orte:
		return "orte"
	case smp:
		return "smp"
	}
	return "unknown"
}

// MarshalText implements the encoding.TextMarshaler interface.
func (p PEName) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func ParseParallelEnvironment(v string) PEName {
	switch v {
	case "orte":
		return orte
	case "smp":
		return smp
	}
	return undefined
}

type ParallelConfig struct {
	Name  PEName
	Slots int
}
type PEFlag struct {
	Value         ParallelConfig
	originalValue string
	isSet         bool
}

func (f *PEFlag) SetValue(s string) error {
	f.originalValue = s

	v := strings.Fields(s)
	if len(v) != 2 {
		return errors.New("incorrect parallel env values")
	}
	penv := ParseParallelEnvironment(v[0])
	if penv == undefined {
		return errors.New("unknown parallel environment")
	}
	slots, err := strconv.Atoi(v[1])
	if err != nil {
		return err
	}

	f.isSet = true
	f.Value = ParallelConfig{
		Name:  penv,
		Slots: slots,
	}

	return nil
}

func (f *PEFlag) OriginalValue() string {
	return f.originalValue
}
func (f *PEFlag) IsSet() bool {
	return f.isSet
}
func (f PEFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Value)
}

func (f *PEFlag) UnMarshalJSON(data []byte) error {
	var pec ParallelConfig
	if err := json.Unmarshal(data, &pec); err != nil {
		return err
	}
	// honestly not sure what an unmarshalled default of pec.Name
	// would be without testing
	if pec.Name == undefined || pec.Slots == 0 {
		f.isSet = false
		f.Value = pec
	}

	return nil
}
