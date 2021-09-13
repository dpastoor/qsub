package parser

import "strings"

// Args are the qsub Arguments
type Args struct {
	Script string
	Flags  map[string]string
}

func ParseArgs(args []string) (Args, error) {
	result := make(map[string]string)
	var script string
	var currentKey string

	totalArgs := len(args)
	for i, arg := range args {
		if i == totalArgs-1 {
			script = arg
			continue
		}
		if strings.HasPrefix(arg, "-") {
			currentKey = strings.TrimPrefix(arg, "-")
			result[currentKey] = ""
			continue
		}
		if result[currentKey] == "" {
			result[currentKey] = arg
		} else {
			result[currentKey] = result[currentKey] + " " + arg
		}
	}
	return Args{
		Flags:  result,
		Script: script,
	}, nil
}
