package main

import (
	"os"
	"path/filepath"
)

const projectName = "crown"

var (
	homeDir      string
	xdgConfigDir string
	logDir       string
	logFile      string
	runDir       string
	runFile      string
	configDir    string
	configFile   string
)

func initConstants() (err error) {
	homeDir, err = os.UserHomeDir()
	if err != nil {
		return
	}

	xdgConfigDir, err = os.UserConfigDir()
	if err != nil {
		return
	}

	logDir = filepath.Join(homeDir, "log")
	logFile = filepath.Join(logDir, "crown.log")
	runDir = filepath.Join(homeDir, "run")
	runFile = filepath.Join(runDir, projectName+".pid")
	configDir = filepath.Join(xdgConfigDir, projectName)
	configFile = filepath.Join(configDir, "config.yaml")

	return
}
