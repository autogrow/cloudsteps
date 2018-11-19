package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/ghodss/yaml"
)

func main() {
	var write bool
	var tmplFile, in string

	flag.BoolVar(&write, "w", false, "write the output to the template file")
	flag.StringVar(&tmplFile, "t", "", "the path to the template file")
	flag.StringVar(&in, "in", "", "file containing the step function definition")
	flag.Parse()

	if tmplFile == "" {
		panic("no template file specified")
	}

	if in == "" {
		panic("no step functions file specified")
	}

	data, err := ioutil.ReadFile(tmplFile)
	if err != nil {
		panic(err)
	}

	var x interface{}
	if err = yaml.Unmarshal(data, &x); err != nil {
		panic(err)
	}

	tmpl, err := gabs.Consume(x)
	if err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile(in)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(data, &x); err != nil {
		panic(err)
	}

	sfn, err := gabs.Consume(x)
	if err != nil {
		panic(err)
	}

	res := gabs.New()
	jsonDefintion := sfn.Path("DefinitionString").String()
	name, _ := sfn.Path("StateMachineName").Data().(string)

	res.Set("AWS::StepFunctions::StateMachine", "Type")
	res.Set(sfn.Data(), "Properties")
	res.Delete("Properties", "DefinitionString")
	res.Set(jsonDefintion, "Properties", "DefinitionString", "Fn::Sub")

	tmpl.Set(res.Data(), "Resources", name)

	data, err = yaml.Marshal(tmpl.Data())
	if err != nil {
		panic(err)
	}

	if !write {
		fmt.Println(string(data))
		os.Exit(0)
	}

	if err := ioutil.WriteFile(tmplFile, data, 0644); err != nil {
		panic(err)
	}
}
