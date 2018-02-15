package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

var db *gorm.DB

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
	app := &App{
		&OrmLinkRepository{*db},
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/short-link")
	{
		v1.POST("/", app.createLink)
		v1.GET("/:slug", app.fetchLink)
	}
	router.Run()
}