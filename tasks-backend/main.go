package main
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	)
	
	type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	}
	
	func loadTasks() ([]Task, error) {
	data, err := ioutil.ReadFile("tasks.json")
	if err != nil {
	return nil, err
	}
	
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
	return nil, err
	}
	return tasks, nil
	}
	
	func getTasks(c *gin.Context) {
	tasks, err := loadTasks()
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to load tasks"})
	return
	}
	c.JSON(http.StatusOK, tasks)
	}
	
	func getTaskByID(c *gin.Context) {
	tasks, err := loadTasks()
	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to load tasks"})
	return
	}
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
	return
	}
	
	for _, task := range tasks {
	if task.ID == id {
	c.JSON(http.StatusOK, task)
	return
	}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
	
	func main() {
	router := gin.Default()
	
	router.Use(func(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if c.Request.Method == "OPTIONS" {
	c.AbortWithStatus(http.StatusNoContent)
	return
	}
	c.Next()
	})
	
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	
	router.Run(":3000")
	}
	