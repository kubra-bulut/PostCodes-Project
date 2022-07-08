package main

import (
	"RecapGorm/config"
)

func main() {
	config.InitDB()
	config.MigrateTables()
}
