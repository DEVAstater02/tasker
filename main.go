/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/DEVAstater02/tasker/cmd"
	"github.com/DEVAstater02/tasker/db"

	"github.com/spf13/viper"
)

func main() {
	// Parsing of configuration parameters
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error parsing config file: ", err)
		os.Exit(1)
	}

	// Initializing DB connection
	dbConfig := &db.DBConfig{
		Username: fmt.Sprintf("%v", viper.Get("app.data.mysql.username")),
		HostUrl:  fmt.Sprintf("%v", viper.Get("app.data.mysql.hosturl")),
		DBPort:   fmt.Sprintf("%v", viper.Get("app.data.mysql.port")),
		Database: fmt.Sprintf("%v", viper.Get("app.data.mysql.database")),
	}

	dbConn, err := dbConfig.InitDB()
	if err != nil {
		fmt.Println("Error initializing DB :", err)
		os.Exit(1)
	}

	cmd.Execute(dbConn)
}
