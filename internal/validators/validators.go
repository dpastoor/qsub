package validators

import (
	"errors"
	"fmt"

	"github.com/metrumresearchgroup/qsub/internal/parser"
)

func IsValidParallelEnvironment(pe string) error {
	_, err := parser.ParseParallelEnvironment(pe)
	if err != nil {

		return fmt.Errorf("invalid parallel environment: %s", err)
	}
	return nil
}

// IsValidJoinStreams checks that the joinstream setup is correct
// per the qsub docs:
// Declares if the standard error stream of the job will be merged with the standard output stream of the job.
// An option argument value of oe directs that the two streams will be merged, intermixed, as standard output.
// An option argument value of eo directs that the two streams will be merged, intermixed, as standard error.
// If the join argument is n or the option is not specified, the two streams will be two separate files.
func IsValidJoinStreams(j string) error {
	switch j {
	case "eo", "oe", "n":
		return nil
	}
	return errors.New("invalid join stream, must be eo, oe, or n")
}

func IsValidName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func IsValidBoolFlag(bf string) error {
	if bf != "" {
		return errors.New("bool flags should not have any arguments")
	}
	return nil
}
