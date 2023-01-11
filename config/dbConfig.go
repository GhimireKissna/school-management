package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connect() {
	viper.SetConfigFile("confit.json")
	viper.ReadInConfig()

	dbUser := viper.GetString("username")
	dbPassword := viper.Get("password")
	dbHost := viper.Get("host")
	dbPort := viper.Get("port")
	dbName := viper.Get("database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", &dbUser, &dbPassword, &dbHost, &dbPort, &dbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
