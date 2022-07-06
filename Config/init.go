package Config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s sslmode=disable",
		appDBHost,
		appDBUserName,
		appDBPassword,
		appDBName)
	var err error
	DB, err = gorm.Open(postgres.open)
}
