package orm

import (
	_ "github.com/go-sql-driver/mysql" // ...
	"github.com/jinzhu/gorm"
	"gitlab.com/jorgechavezrnd/go_rest/config"
)

var db *gorm.DB

// CreateConnection ...
func CreateConnection() {
	url := config.URLDatabase()
	if connection, err := gorm.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

// CloseConnection ...
func CloseConnection() {
	db.Close()
}

// CreateTables ...
func CreateTables() {
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
}
