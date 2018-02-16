package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/spf13/viper"
	"github.com/koschos/gols/generators"
	"github.com/koschos/gols/domain"
	"github.com/koschos/gols/app"
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
	db.AutoMigrate(&domain.LinkModel{})
}

func main() {
	application := createApp()
	router := gin.Default()

	v1 := router.Group("/api/v1/short-link")
	{
		v1.POST("/", application.CreateLink)
		v1.GET("/:slug", application.FetchLink)
	}

	router.Run(fmt.Sprintf(":%d", config.port))
}

func createApp() *app.App {
	return &app.App{
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