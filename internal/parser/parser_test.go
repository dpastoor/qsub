package parser

import (
	"os"
	"testing"

	"github.com/metrumresearchgroup/wrapt"
)

// renaming inbound t to tt
func TestParser(tt *testing.T) {
	testDataContents := Script{
		Path:     "testdata/run.sh",
		Contents: "#!/bin/bash\n\n#$ -wd /data/Projects/Misc/ForestPlot/bbr-nonmem-poppk-foce/model/pk/100\n\n/data/apps/bbi nonmem run local 100.ctl\n",
		Md5:      "8534d9fb52cdd03fffd6c8c3e4b9e902",
	}
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
				param: []string{"-cwd", "-pe", "orte", "4", "-V", "testdata/run.sh"},
			},
			want: Args{
				Flags:  map[string]string{"cwd": "", "pe": "orte 4", "V": ""},
				Script: testDataContents,
			},
			wanterr: false,
		},
		{
			name: "param test",
			args: args{
				// this would be what it would look like if -pe "orte 4" with quotes was used as args
				param: []string{"-cwd", "-pe", "orte 4", "-V", "testdata/run.sh"},
			},
			want: Args{
				Flags:  map[string]string{"cwd": "", "pe": "orte 4", "V": ""},
				Script: testDataContents,
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
			wd, _ := os.Getwd()
			got, err := ParseArgs(test.args.param)
			t.A.WantError(test.wanterr, err)
			t.A.Equal(test.want, got)
			t.Log("working directory: ")
			t.Log(wd)
		})
	}
}
