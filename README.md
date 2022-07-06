# GORM
## Getting Started with Project from Scratch

**Creating a Go Module**    
`go mod init(project_name):` This command will create an initial go.mod file to describe your project. 
The file will contain the project’s name and other module dependencies.

**Installing GORM**     
`go get gorm.io/gorm`:Run this code in your terminal to install     
`go get gorm.io/driver/postgres`:Run this code to get postgresql


`go mod tidy`:This command will basically match the go.mod file with the dependencies required in the source files.

* Download all the dependencies that are required in your source files and update go.mod file with that dependency.
* Remove all dependencies from the go.mod file which are not required in the source files.
