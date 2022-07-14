# GORM
## Getting Started with Project from Scratch

**Creating a Go Module**    
`go mod init(project_name):` This command will create an initial go.mod file to describe your project. 
The file will contain the project’s name and other module dependencies.

**Installing GORM**     
`go get gorm.io/gorm`:Run this code in your terminal to install     
`go get gorm.io/driver/postgres`:Run this code to get postgresql


`go mod tidy`:This command will basically match the _go.mod_ file with the dependencies required in the source files.

* Download all the dependencies that are required in your source files and update go.mod file with that dependency.
* Remove all dependencies from the _go.mod_ file which are not required in the source files.


`go mod vendor`: The go mod vendor command constructs a directory named vendor in the main module’s root directory that contains copies of all packages needed to support builds and tests of packages in the main module.

[https://gorm.io/docs/index.html](https://gorm.io/docs/index.html): Document for the PostCodes_test.go

**Excelize**: The goal of Excelize is to create and maintain a Go language version of the Excel Document API to handle xlsx files that conform to the Office Open XML (OOXML) standard. 
With Excelize you can use Go to read and write MS Excel files.
[Link for a detailed information](https://xuri.me/excelize/en/)

`go get github.com/xuri/excelize/v2`: Command for installing excelize package

**Resty**: Simple HTTP and REST client library for Go (inspired by Ruby rest-client)

`go get github.com/go-resty/resty/v2`: Run this command to install resty package

