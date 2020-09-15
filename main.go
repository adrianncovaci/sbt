package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sbt/database"
	"sbt/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

var customItems []models.CustomItem
var db = database.InitDB()

func basic(filename string) {
	filename = fmt.Sprintf("/home/adrian/Downloads/%s", filename)
	audit := models.Audit{Name: filename}
	if err := db.Create(&audit).Error; err != nil {
		fmt.Println(err)
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal(content, &customItems)
	if err2 != nil {
		fmt.Println(err2)
	}
	for _, item := range customItems {
		item.AuditID = audit.Id
		if err = db.Create(&item).Error; err != nil {
			fmt.Println(err)
		}
	}
}

func getAllPolicies() *gorm.DB {
	result := db.Find(&[]models.CustomItem{})
	return result
}

func exportPolicies(content string, fileName string) {
	fileName = fmt.Sprintf("%s.json", fileName)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1280,
		Height: 530,
		Title:  "sbt",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(getAllPolicies)
	app.Bind(exportPolicies)
	app.Run()
}
