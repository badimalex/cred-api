package api

import (
	"cred-api/internal/db"
	"cred-api/internal/logging"
	"cred-api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/people", GetPeople)
	r.POST("/people", CreatePerson)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)
}

func CreatePerson(c *gin.Context) {
	var person services.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := services.EnrichPersonData(&person); err != nil {
		logging.Log.Errorf("Failed to enrich person data: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := db.CreatePerson(&person); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, person)
}

func GetPeople(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	name := c.Query("name")
	surname := c.Query("surname")
	gender := c.Query("gender")
	nationality := c.Query("nationality")
	ageMin, _ := strconv.Atoi(c.Query("age_min"))
	ageMax, _ := strconv.Atoi(c.Query("age_max"))

	people, err := db.GetPeople(page, perPage, name, surname, gender, nationality, ageMin, ageMax)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, people)
}

func UpdatePerson(c *gin.Context) {
	var person services.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := db.UpdatePerson(id, &person); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, person)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	if err := db.DeletePerson(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Person deleted successfully"})
}
