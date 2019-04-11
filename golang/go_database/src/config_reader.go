package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Read YAML config and get variables
func readconfig(file string) (*DBParams, error) {
	var db DBParams
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &db)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
		return nil, err
	}
	return &db, nil

}
