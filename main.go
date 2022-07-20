package main

import (
	"RecapGorm/config"
	"RecapGorm/libs"
)

func main() {
	config.InitDB()
	config.MigrateTables()
	//libs.DownloadFile("pk_list.zip", "https://postakodu.ptt.gov.tr/Dosyalar/pk_list.zip")
	//libs.GetSourceToUnzip("pk_list.zip", "")
	//
	//libs.RemoveFile("pk_list.zip")
	//libs.ChangeFileName("pk_20220413.xlsx", "pk_20220417.xlsx")
	//
	//libs.DownloadFile("pk_list.zip", "https://postakodu.ptt.gov.tr/Dosyalar/pk_list.zip")
	//libs.GetSourceToUnzip("pk_list.zip", "")
	//libs.IsFileEqual("pk_20220417.xlsx", "pk_20220413.xlsx")
	//libs.ChangeFileName("pk_20220413.xlsx", "pk_20220417.xlsx")

	libs.ImportFromExcel("pk_20220417.xlsx")
}
