package controller

import (
	"RecapGorm/config"
	"RecapGorm/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

// GETPostCodesCity gives all cities
func GETPostCodesCity(c *gin.Context) {
	config.InitDB()
	var cities []models.Cities
	config.DB.Table("post_codes").Debug().Distinct("city").Find(&cities)

	if cities == nil {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
	c.JSON(http.StatusOK, cities)
}

// GETPostCodesCitiesCounty gives you the county list of the given city
func GETPostCodesCitiesCounty(c *gin.Context) {
	config.InitDB()
	var counties []models.Counties
	city := c.Param("city")
	config.DB.Table("post_codes").Debug().Where("city = ?", city).Distinct("county").Find(&counties)
	if counties == nil {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
	}
	c.JSON(http.StatusOK, gin.H{"city": city})
	c.JSON(http.StatusOK, counties)
}

// GETPostCodesCountiesTown gives you the town list of the given city and county
func GETPostCodesCountiesTown(c *gin.Context) {
	config.InitDB()
	var towns []models.Towns

	city := c.Param("city")
	county := c.Param("county")

	config.DB.Table("post_codes").Debug().Where("city = ? AND county = ?", city, county).Distinct("town").Find(&towns)

	c.JSON(http.StatusOK, gin.H{"city": city})
	c.JSON(http.StatusOK, gin.H{"county": county})
	c.JSON(http.StatusOK, towns)

}

// GETPostCodesTownsDistrict gives you the district list of the given city, county and town
func GETPostCodesTownsDistrict(c *gin.Context) {
	config.InitDB()
	var districts []models.Districts

	county := c.Param("county")
	town := c.Param("town")

	config.DB.Table("post_codes").Debug().Where("county=? AND town = ?", county, town).Distinct("district").Find(&districts)

	//c.JSON(http.StatusOK, gin.H{"city": city})
	//c.JSON(http.StatusOK, gin.H{"county": county})
	c.JSON(http.StatusOK, gin.H{"town": town})
	c.JSON(http.StatusOK, districts)

}

func GETPostCodesWithDistrict(c *gin.Context) {
	config.InitDB()
	var postCode models.Codes
	town := c.Param("town")
	district := c.Param("district")
	config.DB.Table("post_codes").Debug().Where("town=? AND district = ?", town, district).Distinct("code").Find(&postCode)
	c.JSON(http.StatusOK, gin.H{"district": district})
	c.JSON(http.StatusOK, postCode)
}
