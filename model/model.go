package model

import (
	"fmt"
	"gin-example/pkg"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db *gorm.DB

//所有需要迁移的表
var tables = map[string]interface{}{
	// "combo":        &Combo{},
	// "combo_type":   &ComboType{},
	// "rantal_order": &RentalOrder{},
}

func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", pkg.MysqlConfig.User, pkg.MysqlConfig.Password, pkg.MysqlConfig.Host, pkg.MysqlConfig.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Info("gorm open mysql err")
	}
	log.Println("mysql connect success!!!")
	logrus.WithFields(logrus.Fields{}).Info("mysql connect success!!!")
	sqlDB, err := db.DB()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Info("db.DB err")
	}

	sqlDB.SetMaxIdleConns(10)

	for _, v := range tables {
		db.AutoMigrate(v)
	}

}
