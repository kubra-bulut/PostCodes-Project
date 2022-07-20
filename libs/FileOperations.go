package libs

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//Downloads the file from the given url

func DownloadFile(filePath, url string) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	fmt.Println(dir)
	// Create a Resty Client
	client := resty.New()

	// HTTP response gets saved into file, similar to curl -o flag
	if _, err := client.R().
		SetOutput(filePath).
		Get(url); err != nil {
		return errors.New("failed downloading the file")
	}

	fmt.Println("File downloaded")
	return nil //check this line later
}

//Gets source and use the UnzipFile func to unzip

func GetSourceToUnzip(source, destination string) error {

	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {

		err := UnzipFile(f, destination)
		if err != nil {
			return err
		}

	}

	return nil

}

//Unzip the given zip file

func UnzipFile(f *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}

	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}

//Checks the given files last modification time

func CheckFileModificationTime(fileName string) {
	fileToCheck, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	//Checking the modified time of the current file
	fmt.Println("File Name:", fileToCheck.Name())             // Base name of the file
	fmt.Println("Last Modified Time:", fileToCheck.ModTime()) // Last modification time

}

//Changes the given files name

func ChangeFileName(originalName, newName string) {
	e := os.Rename(originalName, newName)
	if e != nil {
		log.Fatal(e)
	}
}

//Compares two file and check if the two file is equal

func IsFileEqual(oldFile, newFile string) {
	if strings.Contains(oldFile, newFile) == true {
		fmt.Println("files are equal")
		RemoveFile(newFile)
		fmt.Println("New file has been removed")
	} else {
		fmt.Println("files are not equal")
		RemoveFile(oldFile)
		fmt.Println("Old file has been removed")
	}
}

//Removes the given file

func RemoveFile(fileName string) {
	e := os.Remove(fileName)
	if e != nil {
		log.Fatal(e)
	}
}
