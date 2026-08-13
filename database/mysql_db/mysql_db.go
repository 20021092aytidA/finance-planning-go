package mysql_db

import (
	"finance-planning-go/config/env_cfg"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Start() error {
	var err error
	dbCfg := struct {
		dbHost string
		dbPort string
		dbName string
		dbUser string
		dbPass string
	}{
		dbHost: env_cfg.Get("DB_HOST"),
		dbPort: env_cfg.Get("DB_PORT"),
		dbName: env_cfg.Get("DB_NAME"),
		dbUser: env_cfg.Get("DB_USER"),
		dbPass: env_cfg.Get("DB_PASS"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCfg.dbUser, dbCfg.dbPass, dbCfg.dbHost, dbCfg.dbPort, dbCfg.dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
