package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	file := flag.String("file", "", "file")
	flag.Parse()

	fileBytes, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(fmt.Sprintf("error reading input file: %s", err))
	}

	var yamlContent interface{}
	if err = yaml.Unmarshal(fileBytes, &yamlContent); err != nil {
		panic(fmt.Sprintf("error unmarshaling into yaml: %s", err))
	}

	cleanedYaml, err := yaml.Marshal(yamlContent)
	if err != nil {
		panic(fmt.Sprintf("error marshaling into yaml: %s", err))
	}

	fmt.Println(string(cleanedYaml))
}
