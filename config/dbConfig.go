package config

import (
	"fmt"
	// "school_management/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error ){
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()

	dbUser := viper.GetString(`username`)
	dbPassword := viper.GetString("password")
	dbHost := viper.GetString("host")
	dbPort := viper.GetInt("port")
	dbName := viper.GetString("database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}


