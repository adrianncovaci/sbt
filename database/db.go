package database

import (
	"fmt"
	"sbt/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	var db *gorm.DB
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/sbt?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.CustomItem{}, &models.Audit{})
	return db
}
