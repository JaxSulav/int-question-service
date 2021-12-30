package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var AuthAddress string
var GrpcPort string
var GwRestPort string

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// Can be done using this method or viper.Get() method
	AuthAddress = viper.GetString("AuthAddress")

	var ok bool
	GrpcPort, ok = viper.Get("GrpcPort").(string)
	GwRestPort, ok = viper.Get("GwRestPort").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
}
