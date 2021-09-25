package flags

import (
	"github.com/metrumresearchgroup/qsub/internal/validators"
)

func NewValidatorMap() map[string]func(string) error {
	return map[string]func(string) error{
		"cwd": validators.IsBoolFlag,
		"pe":  validators.IsParallelEnvironment,
		"N":   validators.IsName,
		"j":   validators.IsJoinStreams,
		"V":   validators.IsBoolFlag,
	}
}
