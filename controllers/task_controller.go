package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)


// GetAllTasks handles the HTTP GET request to retrieve all tasks.
func GetAllTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch tasks: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

// GetTaskByID handles the HTTP GET request to retrive a task by ID.
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	// check for id value
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id can not be empty",
		})
		return
	}

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("failed to retrieve task: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

// DeleteTask handles the HTTP DELETE request to remove a task by ID.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	// check for id value
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id can not be empty",
		})
		return
	}

	// delete the task
	if err := data.DeleteTaskByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("failed to delete task: %s", err.Error()),
		})
		return
	}

	// send the success message
	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})

}

// UpdateTask handles the HTTP PUT request to fully replace a task with a new one.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	// check for id value
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id can not be empty",
		})
		return
	}

	var body models.Task

	// read the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	// convert the status into lowercase
	body.Status = strings.ToLower(body.Status)

	// check for valid struct
	if err := isValidTask(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update the task
	err := data.UpdateTask(id, body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated successfully",
	})
}

// AddTask handles the HTTP POST request to insert a task into task collections.
func AddTask(c *gin.Context) {
	var body models.Task

	// read the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// convert the status into lowercase
	body.Status = strings.ToLower(body.Status)

	// check for valid struct
	if err := isValidTask(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	new_task, err := data.AddTask(body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "task created successfully",
		"data":    new_task,
	})

}
