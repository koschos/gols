package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/spf13/viper"
	"github.com/koschos/gols/generators"
	"github.com/koschos/gols/domain"
	"github.com/koschos/gols/handlers"
	"github.com/koschos/gols/storage"
)

const (
	configName = "config"
	slugLength = 6
	slugCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	db *gorm.DB
    config *Config
)

func init() {
	//open a db connection
	bootstrap := Bootstrap{viper.New(), configName}
	config = bootstrap.ReadConfig()
	db = bootstrap.createDb(config)

	//Migrate the schema
	db.AutoMigrate(&domain.LinkModel{})
}

func main() {
	defer db.Close()

	slugGenerator := &generators.RandomSlugGenerator{slugLength, slugCharset}
	hashGenerator := &generators.Md5HashGenerator{}
	repository := &storage.GormLinkRepository{*db}

	router := gin.Default()

	router.GET("/redirect/:slug", handlers.RedirectHandler(repository))

	api := router.Group("/api/short-link")
	{
		api.POST("/", handlers.CreateLinkHandler(hashGenerator, slugGenerator, repository))
	}

	router.Run(fmt.Sprintf(":%d", config.port))
}
