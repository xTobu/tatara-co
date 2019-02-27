package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql for gorm
	"github.com/xTobu/tatara-co/server/api/packages/setting"
)

var db *gorm.DB

// Setup Models
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s:3305)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
