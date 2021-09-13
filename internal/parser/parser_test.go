package parser

import (
	"testing"

	"github.com/metrumresearchgroup/wrapt"
)

// renaming inbound t to tt
func TestParser(tt *testing.T) {
	type args struct {
		param []string
	}
	tests := []struct {
		name    string
		args    args
		want    Args
		wanterr bool
	}{
		{
			name: "param test",
			args: args{
				param: []string{"-cwd", "-pe", "orte", "4", "-V", "run.sh"},
			},
			want: Args{
				Flags:  map[string]string{"cwd": "", "pe": "orte 4", "V": ""},
				Script: "run.sh",
			},
			wanterr: false,
		},
		{
			name: "param test",
			args: args{
				// this would be what it would look like if -pe "orte 4" with quotes was used as args
				param: []string{"-cwd", "-pe", "orte 4", "-V", "run.sh"},
			},
			want: Args{
				Flags:  map[string]string{"cwd": "", "pe": "orte 4", "V": ""},
				Script: "run.sh",
			},
			wanterr: false,
		},
	}
	// rename the default tt to 'test'
	for _, test := range tests {
		// also renaming it in the sub-test
		tt.Run(test.name, func(tt *testing.T) {
			// setting a local t for comfort
			// Note, you can '.' import wrapt to get
			// rid of the stutter of wrapt.WrapT.
			t := wrapt.WrapT(tt)

			got, err := ParseArgs(test.args.param)
			t.A.WantError(test.wanterr, err)
			t.A.Equal(test.want, got)
		})
	}
}
