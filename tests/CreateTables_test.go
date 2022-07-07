package tests

import (
	"RecapGorm/config"
	"testing"
)

func TestCreateTable(t *testing.T) {
	config.InitDB()
	config.MigrateTables()
}
