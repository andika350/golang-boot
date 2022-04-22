package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api-project/models"
	"rest-api-project/views"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


type controllerPerson struct {
	db *gorm.DB
}

func NewControllerPerson(db *gorm.DB) *controllerPerson {
	return &controllerPerson{
		db: db,
	}
}

//get  all person data
func (in *controllerPerson) GetPerson(c *gin.Context) {
	var (
		person []models.Person
		result gin.H
	)

	in.db.Find(&person)

	if len(person) == 0 {
		result = gin.H{
			"result": nil,
			"count":0,
		}
	} else {
		var personViews []views.PersonView
		for _, p := range person {
			personView := views.PersonView {
				ID:		int(p.ID),
				FirstName:	p.First_Name,
				LastName: p.Last_Name,
			}

			personViews = append(personViews, personView)
		}
		result = gin.H {
			"result": personViews,
			"count": len(personViews),
		}
	}
	
	c.JSON(http.StatusOK, result)
}


//create person
func (in *controllerPerson) CreatePerson(c *gin.Context) {
	var person models.Person
	
	err := json.NewDecoder(c.Request.Body).Decode(&person)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H {
		"data": person,
	})
}

//delete person by id
func(in *controllerPerson) DeletePersonByID(c *gin.Context) {
	var (
		person models.Person
	)

	id := c.Param("id")

	err := in.db.First(&person, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Delete(&person).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete data success",
	})
}

//update person by id
func (in *controllerPerson) UpdatePersonByID(c *gin.Context) {
	var (
		person 	models.Person
		newPerson models.Person
	)

	id := c.Param("id")

	err := in.db.First(&person, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newPerson)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Model(&person).Updates(newPerson).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "update data success !",
		"data": newPerson,
	})

}