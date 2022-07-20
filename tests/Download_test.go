package tests

import (
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

	// HTTP response gets saved into file, similar to curl -o flag
	if _, err := client.R().
		SetOutput("pk_list.zip").
		Get("https://postakodu.ptt.gov.tr/Dosyalar/pk_list.zip"); err != nil {
		fmt.Println(err.Error())
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
