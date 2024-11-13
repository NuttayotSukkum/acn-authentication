package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configs struct {
	App      AppConfigs
	Database ConfigSQL
}

type AppConfigs struct {
	Port       string
	HmacSecret []byte
}

type ConfigSQL struct {
	DatabaseURL  string
	DatabaseName string
	DatabaseHost string
	DatabaseUser string
	DatabasePass string
}

func initConfigFile() *Configs {
	var cfg Configs
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func InitConfig() *Configs {
	cfg := initConfigFile()
	cfg.App.Port = viper.GetString("App.Port")
	cfg.Database.DatabaseURL = viper.GetString("Database.url")
	cfg.Database.DatabaseName = viper.GetString("Database.dbname")
	cfg.Database.DatabaseHost = viper.GetString("Database.host")
	cfg.Database.DatabaseUser = viper.GetString("Database.username")
	cfg.Database.DatabasePass = viper.GetString("Database.password")
	strHmacSecret := viper.GetString("App.SECRET. hmacSampleSecret")
	cfg.App.HmacSecret = []byte(strHmacSecret)
	fmt.Sprintf("APP Port: %v", cfg.App.Port)
	fmt.Sprintf("APP Port: %v", cfg.Database.DatabaseURL)
	return cfg
}
