package controller

import (
	"RecapGorm/config"
	"RecapGorm/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//GETPostCodes gets all data from database
func GETPostCodes(c *gin.Context) {
	c.JSON(http.StatusOK, config.PostCodes)
}

// GETPostCodesWithID gets data by given id
func GETPostCodesWithID(c *gin.Context) {
	var returnValue models.PostCode
	id_ := c.Param("id")
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id, id_)

	for _, postCode := range config.PostCodes {
		if postCode.ID == id {
			returnValue = postCode
		}
	}
	if returnValue.ID == 0 {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
	c.JSON(http.StatusOK, returnValue)
}

func GETPostCodesWithCity(c *gin.Context) {
	var returnValue models.PostCode

	for _, postCode := range config.PostCodes {
		city := c.Param("city")
		if postCode.City == city {
			returnValue = postCode
			c.JSON(http.StatusOK, returnValue)
		}

	}
	if returnValue.ID == 0 {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
}

func GETPostCodesWithCounty(c *gin.Context) {
	var returnValue models.PostCode

	for _, postCode := range config.PostCodes {
		county := c.Param("county")
		if postCode.County == county {
			returnValue = postCode
			c.JSON(http.StatusOK, returnValue)
		}

	}
	if returnValue.ID == 0 {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
}
func GETPostCodesWithTown(c *gin.Context) {
	var returnValue models.PostCode

	for _, postCode := range config.PostCodes {
		town := c.Param("town")
		if postCode.Town == town {
			returnValue = postCode
			c.JSON(http.StatusOK, returnValue)
		}

	}
	if returnValue.ID == 0 {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
}
func GETPostCodesWithDistrict(c *gin.Context) {
	var returnValue models.PostCode

	for _, postCode := range config.PostCodes {
		district := c.Param("district")
		if postCode.District == district {
			returnValue = postCode
			c.JSON(http.StatusOK, returnValue)
		}

	}
	if returnValue.ID == 0 {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
}
