/*
Copyright © 2021 Metrum Research Group <developers@metrumrg.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"qsub/internal/flags"
	"qsub/internal/parser"
	"reflect"

	"github.com/fatih/structs"
	"github.com/fatih/structtag"
)

func main() {
	res, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	flagSet := &flags.FlagSet{}
	sfs := structs.New(flagSet)
	names := sfs.Names()
	// TODO: should provide mode to enforce strict flags where only flags set
	// should be those understandable to qsub
	for _, name := range names {
		fld := sfs.Field(name)
		// gotta dereference here
		tag, _ := reflect.TypeOf(*flagSet).FieldByName(fld.Name())
		// ... and start using structtag by parsing the tag
		tags, err := structtag.Parse(string(tag.Tag))
		if err != nil {
			panic(err)
		}
		flg, err := tags.Get("flags")
		if err != nil {
			panic(err)
		}
		val, ok := res.Flags[flg.Name]
		if !ok {
			// flag not set
			continue
		}
		switch tag.Type.String() {
		case "flags.PEFlag":
			v := flags.PEFlag{}
			err = v.SetValue(val)
			if err != nil {
				fmt.Printf("could not set parallel environment `%s` \n", val)
				log.Fatal(err)
			}
			err = fld.Set(v)
			if err != nil {
				fmt.Println("could not set underlying struct")
				log.Fatal(err)
			}
		case "flags.StringFlag":
			v := flags.StringFlag{}
			err = v.SetValue(val)
			if err != nil {
				fmt.Println("could not set StringFlag")
				log.Fatal(err)
			}
			err = fld.Set(v)
			if err != nil {
				fmt.Println("could not set underlying struct")
				log.Fatal(err)
			}
		case "flags.BoolFlag":
			v := flags.BoolFlag{}
			err = v.SetValue(val)
			if err != nil {
				fmt.Println("could not set BoolFlag")
				log.Fatal(err)
			}
			err = fld.Set(v)
			if err != nil {
				fmt.Println("could not set underlying struct")
				log.Fatal(err)
			}
		}

	}
	fmt.Println(prettyJson(flagSet))
	//fmt.Println(prettyJson(res))
}

func prettyJson(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}
