package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		MongoDB struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:'mongodb"`
	} `json:"database"`
}

func Loader() *Config {
	var conf Config

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// v := viper.AllSettings()
	conf.Database.MongoDB.URL = viper.GetString("database.mongodb.url")
	conf.Database.MongoDB.Name = viper.GetString("database.mongodb.name")

	// fmt.Println(conf)

	return &conf
}
