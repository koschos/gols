package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/koschos/gols/generators"
	"github.com/spf13/viper"
)

var db *gorm.DB
var config *Config

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	//open a db connection
	reader := ConfigReader{viper.New(), "config"}
	config = reader.Read()

	db = createDb(config)

	//Migrate the schema
	db.AutoMigrate(&linkModel{})
}

func main() {
	app := createApp()
	router := gin.Default()

	v1 := router.Group("/api/v1/short-link")
	{
		v1.POST("/", app.createLink)
		v1.GET("/:slug", app.fetchLink)
	}

	router.Run(fmt.Sprintf(":%d", config.port))
}

func createApp() *App {
	return &App{
		&OrmLinkRepository{*db},
		&generators.RandomSlugGenerator{6, charset},
		&generators.Md5HashGenerator{},
	}
}

func createDb(config *Config) *gorm.DB {
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