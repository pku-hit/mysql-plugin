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
	Protocol string `yaml:"protocol"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Param    string `yaml:"param"`
}{}

var Db *gorm.DB

func init() {
	config, errRd := ioutil.ReadFile("config/database.yml")
	if errRd != nil {
		fmt.Print(errRd)
	}
	yaml.Unmarshal(config, &Database)

	var err error
	Db, err = gorm.Open(Database.Type, fmt.Sprintf("%s:%s@%s(%s)/%s?%s",
		Database.User,
		Database.Password,
		Database.Protocol,
		Database.Host,
		Database.Name,
		Database.Param))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer Db.Close()
}
