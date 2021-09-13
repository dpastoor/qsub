package flags

import "encoding/json"

type Flag interface {
	IsSet() bool
	OriginalValue() string
	SetValue(string) error
	Type() string
}

type FlagSet struct {
	CurrentWorkingDirectory BoolFlag   `json:",omitempty" flags:"cwd"`
	PropogateEnvVars        BoolFlag   `json:",omitempty" flags:"V"`
	ParallelEnvironment     PEFlag     `json:",omitempty" flags:"pe"`
	Name                    StringFlag `json:",omitempty" flags:"N"`
	// The option -j oe will merge stderr into stdout (and hence the -e option does not make sense),
	// the option -j eo will merge stdout into stderr.
	// TODO: if using this, prob worth making another custom flag type
	JoinStreams StringFlag `json:",omitempty" flags:"j"`
}

func (f FlagSet) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		CurrentWorkingDirectory bool   `json:",omitempty" flags:"cwd"`
		PropogateEnvVars        bool   `json:",omitempty" flags:"V"`
		ParallelEnvironment     PEFlag `json:",omitempty" flags:"pe"`
		Name                    string `json:",omitempty" flags:"N"`
		// The option -j oe will merge stderr into stdout (and hence the -e option does not make sense),
		// the option -j eo will merge stdout into stderr.
		// TODO: if using this, prob worth making another custom flag type
		JoinStreams string `json:",omitempty" flags:"j"`
	}{
		CurrentWorkingDirectory: f.CurrentWorkingDirectory.IsSet(),
		PropogateEnvVars:        f.PropogateEnvVars.IsSet(),
		ParallelEnvironment:     f.ParallelEnvironment,
		Name:                    f.Name.Value,
		JoinStreams:             f.JoinStreams.Value,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}
