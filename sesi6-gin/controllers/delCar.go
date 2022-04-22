package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DelCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
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

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("car with id %v has been succesfully deleted", carID),
	})

}