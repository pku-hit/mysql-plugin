package mysql_plugin

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Database = struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}{}

var db *gorm.DB

func init() {
	config, errRd := ioutil.ReadFile("config/database.yml")
	if errRd != nil {
		fmt.Print(errRd)
	}
	yaml.Unmarshal(config, &Database)

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
