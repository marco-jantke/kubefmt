package main

import (
	"flag"
	"fmt"

	"io"
	"os"

	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	file := flag.String("file", "", "file")
	flag.Parse()

	var input io.Reader = os.Stdin
	var err error

	if *file != "" {
		input, err = os.Open(*file)
		if err != nil {
			panic(fmt.Sprintf("error opening input file: %s", err))
		}
	}

	inputYaml, err := ioutil.ReadAll(input)
	if err != nil {
		panic(fmt.Sprintf("error reading yaml content: %s", err))
	}

	var yamlContent interface{}
	if err = yaml.Unmarshal(inputYaml, &yamlContent); err != nil {
		panic(fmt.Sprintf("error unmarshaling into yaml: %s", err))
	}

	cleanedYaml, err := yaml.Marshal(yamlContent)
	if err != nil {
		panic(fmt.Sprintf("error marshaling into yaml: %s", err))
	}

	fmt.Println(string(cleanedYaml))
}
