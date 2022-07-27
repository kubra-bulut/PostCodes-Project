package rooter

import (
	"RecapGorm/controller"
	"github.com/gin-gonic/gin"
)

func PostCodesRoot(api *gin.RouterGroup) {
	api.GET("/postCodes", controller.GETPostCodes)
	api.GET("/postCodes/:id", controller.GETPostCodesWithID)
	//api.GET("/postCodes/:city", controller.GETPostCodesWithCity)
	//api.GET("/postCodes/:county", controller.GETPostCodesWithCounty)
	//api.GET("/postCodes/:town", controller.GETPostCodesWithTown)
	api.GET("/district/:district", controller.GETPostCodesWithDistrict)
}
