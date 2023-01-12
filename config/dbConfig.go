package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var db *gorm.DB

func Connect() (*gorm.DB,error ){
	viper.SetConfigFile("confit.json")
	viper.ReadInConfig()

	dbUser := viper.GetString("username")
	dbPassword := viper.Get("password")
	dbHost := viper.Get("host")
	dbPort := viper.Get("port")
	dbName := viper.Get("database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}
	return db, err
}

