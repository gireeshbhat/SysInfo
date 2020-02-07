package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"errors"
	"fmt"
)

type Config struct {
	DatabaseName		string `json:"database"`
	DatabaseHost		string `json:"host"`
	DatabasePort		string `json:"port"`
	UserName			string `json:"user-name"`
	Password			string `json:"password"`
}

var configData *Config

func Set(rootPath, configFileName string) error {
	configFilePath := filepath.Join(rootPath, configFileName)
	if configData != nil {
		return nil
	}
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		return errors.New(err.Error())
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			fmt.Println("Error closing config file: ", err.Error())
		}
	}()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return errors.New(err.Error())
	}
	var config Config
	err = json.Unmarshal([]byte(data), &config)
	configData = &config
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

