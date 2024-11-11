package controllers

import (
	"net/http"
	"swagger/database"
	"swagger/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @GetAllCards godoc
// @Summary Get all cars
// @Description Get all cards
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()
	var cars []models.Car

	err := db.Find(&cars).Error

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// @CreateCars godoc
// @Summary Create a new car entry.
// @Description Create a new car entry.
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "Input Create Car"
// @Router /cars [post]
func CreateCars(c *gin.Context) {
	var db = database.GetDB()
	var car models.Car

	userData := c.MustGet("userData").(jwt.MapClaims)

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// get user id
	userID := int(userData["id"].(float64))

	// input data
	carInput := models.Car{
		Merk:     car.Merk,
		TypeCars: car.TypeCars,
		Pemilik:  car.Pemilik,
		Harga:    car.Harga,
	}
	carInput.UserID = userID

	// proses create
	err := db.Create(&carInput).Error
	// error handling
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	// response success
	c.JSON(http.StatusOK, gin.H{
		"data": carInput,
	})
}

// @UpdateCars godoc
// @Summary Update car identified by the given ID
// @Description Update the card corresponding to the input ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 20 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCars(c *gin.Context) {
	var db = database.GetDB()
	var car models.Car

	err := db.First(&car, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input

	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

// @DeleteCars godoc
// @Summary
func DeleteCars(c *gin.Context) {
	var db = database.GetDB()
	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
