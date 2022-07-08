package tests

import (
	"RecapGorm/config"
	"fmt"
	"testing"
)

func TestImportFromExcel(t *testing.T) {
	config.InitDB()
	f, err := excelize.OpenFile("C:\\Users\\K\\GolandProjects\\RecapGorm\\pk_20220413.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

}
