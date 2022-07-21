package libs

import (
	"fmt"
)

func DoTheOperations() {

	err := DownloadFile("pk_list.zip", "https://postakodu.ptt.gov.tr/Dosyalar/pk_list.zip")
	if err != nil {
		fmt.Println(err)
	}
	err = GetSourceToUnzip("pk_list.zip", "download")
	if err != nil {
		fmt.Println(err)
	}

	FindUpdatedFile()

	ImportFromExcel(secondFile)

}
