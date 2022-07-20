package libs

import (
	"RecapGorm/config"
	"RecapGorm/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func FindDataWithId(id int) {
	config.InitDB()
	var postCode models.PostCode
	if err := config.DB.First(&postCode, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Kayıt bulunamadı durumunda yazılacak kod bu kısımda olacak")
	}
	PrintPrettyStruct(postCode)
}

func GetAllDataWithGivenCode(code string) {
	config.InitDB()
	type Result struct {
		Code     string `json:"code"`
		City     string `json:"city"`
		County   string `json:"county"`
		Town     string `json:"town"`
		District string `json:"district"`
	}
	var statistics []Result
	config.DB.Model(&models.PostCode{}).
		Where("code", code).
		Debug().Find(&statistics)
	PrintPrettyStruct(statistics)

}

func GetAllDataWithGivenCity(city string) {
	config.InitDB()
	type Result struct {
		Code     string `json:"code"`
		City     string `json:"city"`
		County   string `json:"county"`
		Town     string `json:"town"`
		District string `json:"district"`
	}
	var statistics []Result
	config.DB.Model(&models.PostCode{}).
		Where("city", city).
		Debug().Find(&statistics)
	PrintPrettyStruct(statistics)

}
func GetAllDataWithGivenTown(town string) {
	config.InitDB()
	type Result struct {
		Code     string `json:"code"`
		City     string `json:"city"`
		County   string `json:"county"`
		Town     string `json:"town"`
		District string `json:"district"`
	}
	var statistics []Result
	config.DB.Model(&models.PostCode{}).
		Where("town", town).
		Debug().Find(&statistics)
	PrintPrettyStruct(statistics)

}
