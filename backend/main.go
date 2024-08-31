package main

import (
	"task_manager/controller"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDB()
	router.POST("/users/:firebase_uid", controller.CreateNewUser)
	router.GET("/users/:firebase_uid/tasks", controller.ListAllTasks)
	router.GET("/users/:firebase_uid/tasks/:task_id", controller.GetTask)
	router.POST("/users/:firebase_uid/tasks", controller.CreateNewTask)
	router.PUT("/users/:firebase_uid/tasks/:task_id", controller.UpdateTask)
	router.DELETE("/users/:firebase_uid/tasks/:task_id", controller.DeleteTask)
	router.Run()
	defer models.DB.Close()

}
