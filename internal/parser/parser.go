package parser

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"
)

// Args are the qsub Arguments
type Script struct {
	Path     string
	Contents string
	Md5      string
}
type Args struct {
	Script Script
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

	fb, err := ioutil.ReadFile(script)
	if err != nil {
		return Args{}, err
	}
	return Args{
		Flags:  result,
		Script: Script{Path: script, Contents: string(fb), Md5: fmt.Sprintf("%x", md5.Sum(fb))},
	}, nil
}
