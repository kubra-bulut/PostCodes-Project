package config

import (
	"RecapGorm/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB            *gorm.DB
	appDBName     = os.Getenv("recapDBName")
	appDBHost     = os.Getenv("gormDBHost")
	appDBUserName = os.Getenv("gormDBUserName")
	appDBPassword = os.Getenv("gormDBPassword")
)

func InitDB() {
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s",
		appDBHost,
		appDBUserName,
		appDBPassword,
		appDBName)
	var err error
	DB, err = gorm.Open(postgres.Open(cnnString), &gorm.Config{})
	if err != nil {
		log.Println("DB Error", err)
	}
	fmt.Println("DB Connected")
}

func MigrateTables() {
	_ = DB.AutoMigrate(&models.PostCode{})
}

func DeleteTable() {
	DB.Migrator().DropTable(&models.PostCode{})
}
