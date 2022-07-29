package main

import (
	"RecapGorm/config"
	"RecapGorm/rooter"
	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()
	r := gin.Default()
	api := r.Group("")
	rooter.PostCodesRoot(api)
	r.Run()
}
