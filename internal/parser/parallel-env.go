package parser

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
