package mysql_plugin

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Database = struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}{}

var db *gorm.DB

func init() {
	configor.Load(&Database, "config/database.yml")

	var err error
	db, err = gorm.Open(Database.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Database.User,
		Database.Password,
		Database.Host,
		Database.Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
