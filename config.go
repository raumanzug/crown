package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type crown_conf_s struct {
	Crontabs []crontab_s `yaml:""`
}

type crontab_s struct {
	Label   string   `yaml:""`
	Spec    string   `yaml:""`
	Command string   `yaml:""`
	Args    []string `yaml:""`
}

func loadConfigFile(actualConfigFile string, pConfigData *crown_conf_s) (err error) {
	in, err := os.ReadFile(actualConfigFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(in, pConfigData)
	if err != nil {
		return
	}

	return
}
