package parser

import (
	"errors"
	"fmt"
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

func checkParallelEnvironment(v string) error {
	switch v {
	case "orte", "smp":
		return nil
	}
	return fmt.Errorf("parallel environment must be either orte or smp, not %s", v)
}

func ParseParallelEnvironment(pe string) (ParallelConfig, error) {

	fields := strings.Fields(pe)
	if len(fields) != 2 {
		return ParallelConfig{}, errors.New("parallel environment should only have two fields, the environment and core count")
	}
	err := checkParallelEnvironment(fields[0])
	if err != nil {
		return ParallelConfig{}, err
	}
	i, err := strconv.Atoi(fields[1])
	if err != nil || i < 1 {
		return ParallelConfig{}, fmt.Errorf("invalid slot configuration - must be an integer > 0, not %s", fields[1])
	}
	return ParallelConfig{
		Name:  fields[0],
		Slots: i,
	}, nil
}

type ParallelConfig struct {
	Name  string
	Slots int
}
