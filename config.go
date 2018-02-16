package main

import "github.com/spf13/viper"

type Config struct {
	port   int
	dbHost string
	dbPort int
	dbUser string
	dbPass string
	dbName string
}

type ConfigReader struct {
	v          *viper.Viper
	configName string
}

func (r *ConfigReader) Read() *Config {
	var v = *r.v
	v.SetConfigName(r.configName)
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return &Config{
		v.GetInt("port"),
		v.GetString("db.host"),
		v.GetInt("db.port"),
		v.GetString("db.user"),
		v.GetString("db.password"),
		v.GetString("db.name"),
	}
}
