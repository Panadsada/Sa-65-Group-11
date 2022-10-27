package controller

 

import (

             "github.com/Panadsada/sa-65-example/entity"

           "github.com/gin-gonic/gin"

           "net/http"
		

)


func ListDoctors(c *gin.Context) {

	var doctors []entity.Doctor

	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&doctors).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}
	c.JSON(http.StatusOK, gin.H{"data": doctors})

}
