package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigData struct {
	DbAdmin      string `json:"DbAdmin"`
	DbPWD        string `json:"DbPWD"`
	DbPath       string `json:"DbPath"`
	DbService    string `json:"DbService"`
	Port         int    `json:"Port"`
	LinesPerPage int    `json:"LinesPerPage"`
}

var FileConfig ConfigData

func InitConfig(path string) (bool, *ConfigData) {

	jsFile, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return false, nil
	}

	defer jsFile.Close()

	byteValue, _ := ioutil.ReadAll(jsFile)
	err = json.Unmarshal([]byte(byteValue), &FileConfig)
	if err != nil {
		log.Print(err)
	}
	return true, &FileConfig
}
