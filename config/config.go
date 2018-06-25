package config

import (
	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"
)

// ConfigVariables Variable for Environment Configuration
type ConfigVariables struct {
	AppEnv                 	   string `json:"app_env"`
	Port                       string `json:"port"`
	DatabaseUrl				   string `json:"database_url"`
	DatabaseName			   string `json:"database_name"`
	CustomerCollection		   string `json:"customer_collection"`
}

func getVariablesByFile() ConfigVariables {
	file, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/henriqueholanda/cache-failover/config/config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	configVariables := ConfigVariables{}
	json.Unmarshal([]byte(file), &configVariables)

	return configVariables
}

//GetConfigs Return the environment configurations
func GetConfigs() ConfigVariables {
	isDev := os.Getenv("APP_ENV") != "live" && os.Getenv("APP_ENV") != "staging"
	if isDev {
		return getVariablesByFile()
	}

	configs := ConfigVariables{}
	configs.AppEnv = os.Getenv("APP_ENV")
	configs.Port = os.Getenv("PORT")
	configs.DatabaseUrl = os.Getenv("DB_URL")
	configs.DatabaseName = os.Getenv("DB_NAME")
	configs.CustomerCollection = os.Getenv("CUSTOMER_COLLECTION")

	return configs
}

