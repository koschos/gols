package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/koschos/gols/generators"
	"crypto/md5"
)

var db *gorm.DB

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:@/gols")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database. %s", err))
	}
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
	router.Run()
}

func createApp() *App {
	return &App{
		&OrmLinkRepository{*db},
		&generators.RandomSlugGenerator{6, charset},
		&generators.Md5HashGenerator{md5.New()},
	}
}