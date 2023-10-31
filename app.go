package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks []Task

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func addTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func deleteTasks(c *gin.Context) {
	taskID := c.Param("id")

	ID, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id da tarefa invalido "})
		return
	}

	for i, task := range tasks {
		if task.ID == ID {
			tasks = append(tasks[:i], tasks[+1:]...)
			c.JSON(http.StatusOK, gin.H{"sucesso": "task excluida"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "tarefa não encontrada"})
}

func main() {
	r := gin.Default()

	tasks = append(tasks, Task{ID: 1, Name: "exemplo de tarefa 1"})
	tasks = append(tasks, Task{ID: 2, Name: "exemplo de tarefa 2"})

	r.DELETE("/tasks/:id", deleteTasks)
	r.GET("/tasks", getTasks)
	r.POST("/tasks", addTask)

	r.Run(":8080")

}
