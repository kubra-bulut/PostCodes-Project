package tests

import (
	"RecapGorm/config"
	"RecapGorm/libs"
	"RecapGorm/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

//Creating Record
func TestPostCodesCreate(t *testing.T) {
	config.InitDB()
	var postCode models.PostCode
	var location models.Location

	location.X = 34.3
	location.Y = 33.2

	postCode.Code = "34087"
	postCode.City = "İstanbul"
	postCode.County = "Fatih"
	postCode.Town = "Balat"
	postCode.District = "Balat"
	postCode.Location = location.ToByteArray()
	postCode.Tags = []string{"minnak", "kurbaa", "filan"}
	config.DB.Create(&postCode)
}

//To efficiently insert large number of records, pass a slice to the Create method.
func TestPostCodesBatchInsert(t *testing.T) {
	config.InitDB()

	//You can do it in both ways.

	//var postCodes = []models.PostCode{
	//	{Code: "34087", City: "İstanbul", County: "Fatih", Town: "Balat", District: "Balat", UploadDate: time.Now()},
	//	{Code: "34088", City: "İstanbul", County: "Fatih", Town: "Balat", District: "Balat", UploadDate: time.Now()},
	//}

	var postCodes []models.PostCode
	postCodes = append(postCodes, models.PostCode{Code: "34087", City: "İstanbul", County: "Fatih", Town: "Balat", District: "Atikali Mah.", UploadDate: time.Now()})
	postCodes = append(postCodes, models.PostCode{Code: "34088", City: "İstanbul", County: "Fatih", Town: "Balat", District: "Balat", UploadDate: time.Now()})
	postCodes = append(postCodes, models.PostCode{Code: "34089", City: "İstanbul", County: "Fatih", Town: "Balat", District: "Dervişali Mah.", UploadDate: time.Now()})
	config.DB.Create(&postCodes)
}

//Creating Hooks
//GORM allows user defined hooks to be implemented for BeforeSave, BeforeCreate, AfterSave, AfterCreate.

//If you want to skip Hooks methods, you can use the SkipHooks session mode.
func TestPostCodesCreateSkipHooks(t *testing.T) {
	config.InitDB()
	var postCode models.PostCode
	postCode.Code = "34087"
	postCode.City = "İstanbul"
	postCode.County = "Fatih"
	postCode.Town = "Balat"
	postCode.District = "Balat"
	result := config.DB.Session(&gorm.Session{SkipHooks: true}).Create(&postCode)
	fmt.Println(result.Error)
}

// Get the first record ordered by primary key
func TestPostCodesFirst(t *testing.T) {
	config.InitDB()
	var postCode models.PostCode
	config.DB.First(&postCode)
	libs.PrintPrettyStruct(postCode)
}

// Get last record, ordered by primary key desc
func TestPostCodesLast(t *testing.T) {
	config.InitDB()
	var postCode models.PostCode
	config.DB.Last(&postCode)
	libs.PrintPrettyStruct(postCode)
}

// SELECT * FROM users WHERE id = 10;
func TestPostCodesFirstWithID(t *testing.T) {
	config.InitDB()
	var postCode models.PostCode
	if err := config.DB.First(&postCode, 2).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Kayıt bulunamadı durumunda yazılacak kod bu kısımda olacak")
	}
	libs.PrintPrettyStruct(postCode)
}

// SELECT * FROM users WHERE id IN (1,2,3);
func TestPostCodesFindWithIDs(t *testing.T) {
	config.InitDB()
	var postCodes []models.PostCode
	config.DB.Debug().Find(&postCodes, []uint64{3, 5, 6})
	libs.PrintPrettyStruct(postCodes)
}

// Get all matched records
func TestPostCodesFindWhere(t *testing.T) {
	config.InitDB()
	var postCodes models.PostCode
	config.DB.Debug().Where("district=? and city=?", "Dervişali Mah.", "İstanbul").Find(&postCodes)
	libs.PrintPrettyStruct(postCodes)
}

//Select allows you to specify the fields that you want to retrieve from database.
//Otherwise, GORM will select all fields by default.
func TestPostCodesFindWithSelectedField(t *testing.T) {
	config.InitDB()
	var postCodes models.PostCode
	config.DB.Debug().Select("code").Find(&postCodes)
	libs.PrintPrettyStruct(postCodes)
}

//Select specific field with a new struct
func TestPostCodesFindWithSelectedFieldWithNewStruct(t *testing.T) {
	config.InitDB()
	type Minnak struct {
		ID   uint   `json:"id"`
		Code string `json:"code"`
	}
	var postCodes []Minnak
	config.DB.Table("post_codes").Debug().Select("id", "code").Or("code desc").Find(&postCodes)
	libs.PrintPrettyStruct(postCodes)
}

func TestPostCodesFindWithSelectedFieldWithGroupByStruct(t *testing.T) {
	config.InitDB()
	type Result struct {
		District    string `json:"district"`
		RecordCount int    `json:"record_count"`
	}
	var statistics []Result
	config.DB.Model(&models.PostCode{}).
		Select("district, count(*) record_count").
		Debug().
		Group("district").
		Find(&statistics)
	libs.PrintPrettyStruct(statistics)
}

func TestPostCodesDeleteRecordWithIDs(t *testing.T) {
	config.InitDB()
	var postCodes models.PostCode
	config.DB.Delete(&postCodes, []int{1, 2, 3, 4, 5, 6})
}
