package config

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Version    bool
	Hidebanner bool
	Watchman   string
	Build      string
	Run        string
}

func LoadConfig(cmd *cobra.Command) Config {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		log.Printf("Failed to bind flags, error: %s\n", err.Error())
	}

	viper.SetEnvPrefix("AUTOFRESH")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigFile("autofresh-config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read config file, error %s\n", err.Error())
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Unable to unmarshal into struct, error: %s\n", err.Error())
	}
	return conf
}
