package tests

import (
	"RecapGorm/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

	// Create a Resty Client
	client := resty.New()

	// Setting output directory path, If directory not exists then resty creates one!
	// This is optional one, if you're planning using absolute path in
	// `Request.SetOutput` and can used together.
	// client.SetOutputDirectory("/Users/jeeva/Downloads")

	// HTTP response gets saved into file, similar to curl -o flag
	if _, err := client.R().
		SetOutput("pk_list.zip").
		Get("https://postakodu.ptt.gov.tr/Dosyalar/pk_list.zip"); err != nil {
		fmt.Println(err.Error())
	}

}

func TestUnzipExcel(t *testing.T) {
	err := models.UnzipSource("pk_list.zip", "C:\\Users\\K\\GolandProjects\\RecapGorm\\tests\\postcodes")
	if err != nil {
		log.Fatal(err)
	}
}

func TestRenameFile(t *testing.T) {
	originalPath := "C:\\Users\\K\\GolandProjects\\RecapGorm\\tests\\postcodes\\pk_20220413.xlsx"
	newPath := "C:\\Users\\K\\GolandProjects\\RecapGorm\\tests\\postcodes\\postcodes.xlsx"
	e := os.Rename(originalPath, newPath)
	if e != nil {
		log.Fatal(e)
	}
}
