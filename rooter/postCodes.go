package rooter

import (
	"RecapGorm/controller"
	"github.com/gin-gonic/gin"
)

func PostCodesRoot(api *gin.RouterGroup) {
	api.GET("/postCodes/cities", controller.GETPostCodesCity)
	api.GET("/city/:city", controller.GETPostCodesCitiesCounty)
	api.GET("/city/:city/county/:county", controller.GETPostCodesCountiesTown)
	api.GET("/county/:county/town/:town", controller.GETPostCodesTownsDistrict)
	api.GET("/town/:town/district/:district", controller.GETPostCodesWithDistrict)
}
