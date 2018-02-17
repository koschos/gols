package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
)

type Config struct {
	port   int
	dbHost string
	dbPort int
	dbUser string
	dbPass string
	dbName string
}

type Bootstrap struct {
	v          *viper.Viper
	configName string
}

func (r *Bootstrap) ReadConfig() *Config {
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

func (r *Bootstrap) createDb(config *Config) *gorm.DB {
	var err error
	var dsn string
	var dbAddress string = ""

	if config.dbHost != "" {
		dbAddress = fmt.Sprintf("%s:%d", config.dbHost, config.dbPort)
	}

	dsn = fmt.Sprintf("%s:%s@%s/%s", config.dbUser, config.dbPass, dbAddress, config.dbName)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database. %s", err))
	}

	return db
}
