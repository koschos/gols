package storage

import (
	"testing"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/koschos/gols/domain"
)

func TestMain(m *testing.M) {
	setupDatabase()
	result := m.Run()
	teardownDatabase()
	os.Exit(result)
}

func setupDatabase() {
	db := createDb()
	db.AutoMigrate(&domain.LinkModel{})
}

func teardownDatabase() {
	db := createDb()
	db.DropTableIfExists(domain.LinkModel{})
}

//
func createDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TEST_NAME"),
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database. %s", err))
	}

	return db
}
