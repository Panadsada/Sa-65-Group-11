package controller

 

import (

             "github.com/Panadsada/sa-65-example/entity"

           "github.com/gin-gonic/gin"

           "net/http"
		   

)


func ListDepartments(c *gin.Context) {

	var departments []entity.Department

	if err := entity.DB().Raw("SELECT * FROM departments").Scan(&departments).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}

	c.JSON(http.StatusOK, gin.H{"data": departments})

}
