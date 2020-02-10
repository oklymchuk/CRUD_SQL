package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ConfigData struct {
	DbAdmin      string `json:"DbAdmin"`
	DbPWD        string `json:"DbPWD"`
	DbPath       string `json:"DbPath"`
	DbService    string `json:"DbService"`
	DbHost       string `json:"DbHost"`
	DbPort       int    `json:"DbPort"`
	Port         int    `json:"Port"`
	LinesPerPage int    `json:"LinesPerPage"`
}

func (conf *ConfigData) InitConfig(pathToConfigFile string) bool {

	jsFile, err := os.Open(pathToConfigFile)
	if err != nil {
		log.Print(err)
		return false
	}

	defer jsFile.Close()

	byteValue, _ := ioutil.ReadAll(jsFile)
	err = json.Unmarshal([]byte(byteValue), conf)
	if err != nil {
		log.Print(err)
	}
	return true
}

func (conf *ConfigData) ConnectString() (dbdriver string, connstr string) {
	cstr := ""
	dbd := ""
	if conf.DbService == "MySQL" {
		cstr = fmt.Sprintf("%s:%s@/%s\n", conf.DbAdmin, conf.DbPWD, conf.DbPath)
		dbd = "mysql"
	} else if conf.DbService == "Postgres" {
		cstr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.DbHost, conf.DbPort, conf.DbAdmin, conf.DbPWD, conf.DbPath)
		dbd = "postgres"
	}
	return dbd, cstr
}
