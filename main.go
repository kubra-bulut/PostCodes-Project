package main

import (
	"RecapGorm/config"
	"RecapGorm/rooter"
	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()
	//config.DeleteTable()
	//libs.DoTheOperations()
	config.DB.Table("post_codes").Find(&config.PostCodes)

	r := gin.Default()
	api := r.Group("")
	rooter.PostCodesRoot(api)
	r.Run()
}
