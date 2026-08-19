package database

import (
	"finance-planning-go/internal/app/config"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	DBType string
}

func (d Database) Start() {
	upperCaseDB := strings.ToUpper(d.DBType)
	var err error
	dbCfg := struct {
		dbHost string
		dbPort string
		dbName string
		dbUser string
		dbPass string
	}{
		dbHost: config.ENV{}.Get(fmt.Sprintf("DB_%s_HOST", upperCaseDB)),
		dbPort: config.ENV{}.Get(fmt.Sprintf("DB_%s_PORT", upperCaseDB)),
		dbName: config.ENV{}.Get(fmt.Sprintf("DB_%s_NAME", upperCaseDB)),
		dbUser: config.ENV{}.Get(fmt.Sprintf("DB_%s_USER", upperCaseDB)),
		dbPass: config.ENV{}.Get(fmt.Sprintf("DB_%s_PASS", upperCaseDB)),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCfg.dbUser, dbCfg.dbPass, dbCfg.dbHost, dbCfg.dbPort, dbCfg.dbName)

	switch strings.ToUpper(d.DBType) {
	case ("MYSQL"):
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("DATABASE CONNECTION FAILED: " + err.Error())
	}
}
