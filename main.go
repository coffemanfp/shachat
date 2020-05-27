package main

import (
	"flag"
	"log"

	"github.com/coffemanfp/shachat/config"
	"github.com/coffemanfp/shachat/databases"
	"github.com/coffemanfp/shachat/router"
)

var (
	configFile    string
	configFileDef string = "config.yaml"
)

func main() {
	r := router.NewRouter()
	r.Run(":8080")
}

func init() {
	initSettings()
	initDatabase()
}

func initSettings() {
	err := config.SetSettingsByFile(configFileDef)
	if err != nil {
		log.Fatalln(err)
	}
}

func initDatabase() {
	_, err := databases.OpenConn()
	if err != nil {
		log.Fatalln(err)
	}
}

func initFlags() {
	flag.StringVar(&configFile, "config-file", configFileDef, "Config file for the server settings.")

	flag.Parse()
}
