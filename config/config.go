package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/models"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("fail to load file")
	}
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectToDB() {
	var dbConfig DBConfig = DBConfig{
		Username: os.Getenv("PRA_DBUSER"),
		Password: os.Getenv("PRA_DBPASSWORD"),
		Host:     os.Getenv("PRA_DBHOST"),
		Port:     os.Getenv("PRA_DBPORT"),
		Name:     os.Getenv("PRA_DBNAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	var err error
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Database Connection Error")
	}

	migration()
}

func migration() {
	DB.AutoMigrate(&models.User{}, &models.Loker{})
}