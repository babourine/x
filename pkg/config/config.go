package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type hasInit interface {
	Init() error
}

// LoadYAML unmarshals yaml file into an object
func LoadYAML(configFile string, obj interface{}) error {
	return loadConfig(configFile, obj, yaml.Unmarshal)
}

// LoadJSON unmarshals json file into an object
func LoadJSON(configFile string, obj interface{}) error {
	return loadConfig(configFile, obj, json.Unmarshal)
}

func loadConfig(configFile string, obj interface{}, fn func([]byte, interface{}) error) error {

	// read full file content
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	// unmarshal object
	if err := fn(content, obj); err != nil {
		return err
	}

	// if object has Init function, call it
	if i, yes := obj.(hasInit); yes {
		return i.Init()
	}

	return nil

}
