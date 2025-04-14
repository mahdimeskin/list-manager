package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mahdimeskin/list-manager/models"
)

type CreateToDoInput struct {
	Title string `json:"title" binding:"required"`
	Completed  bool   `json:"completed"`
	Note string `json:"note"`
}

type UpdateToDoInput struct {
	Title string `json:"title"`
	Completed  bool   `json:"completed"`
	Note string `json:"note"`
}

func GetToDos(c *gin.Context){
	var todos []models.ToDo
	models.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func GetToDo(c *gin.Context){
	var todo models.ToDo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func CreateToDo(c *gin.Context){
	var input CreateToDoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.ToDo{Title: input.Title, Completed: input.Completed, Note: input.Note}
	models.DB.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func UpdateToDo(c *gin.Context) {
	var todo models.ToDo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateToDoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todo).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteToDo(c *gin.Context){
	var todo models.ToDo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

