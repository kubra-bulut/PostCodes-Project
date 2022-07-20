package tests

import (
	"RecapGorm/config"
	"RecapGorm/models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
	"testing"
)

func TestImportFromExcel(t *testing.T) {
	config.InitDB()
	f, err := excelize.OpenFile("postcodes.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var postCodes []models.PostCode
	for i, row := range rows {
		if i == 0 {
			continue
		}
		postCodes = append(postCodes, models.PostCode{
			City:     strings.TrimSpace(row[0]),
			County:   strings.TrimSpace(row[1]),
			Town:     strings.TrimSpace(row[2]),
			District: strings.TrimSpace(row[3]),
			Code:     strings.TrimSpace(row[4]),
		})
		if len(postCodes) == 1000 {
			config.DB.Create(&postCodes)
			postCodes = []models.PostCode{}
			fmt.Println("1000 records inserted")
		}
	}
	if len(postCodes) > 0 {
		config.DB.Create(&postCodes)
		fmt.Println("Remaining records inserted")
	}
}
