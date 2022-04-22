package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H {
		"car": carData,
	})

}