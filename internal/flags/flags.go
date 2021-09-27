package flags

import (
	"github.com/metrumresearchgroup/qsub/internal/validators"
)

func NewValidatorMap() map[string]func(string) error {
	return map[string]func(string) error{
		"cwd": validators.IsValidBoolFlag,
		"pe":  validators.IsValidParallelEnvironment,
		"N":   validators.IsValidName,
		"j":   validators.IsValidJoinStreams,
		"V":   validators.IsValidBoolFlag,
	}
}
