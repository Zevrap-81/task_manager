package controller

import (
	"net/http"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func CreateNewUser(c *gin.Context) {
	firebase_uid := c.Param("firebase_uid")
	username := c.PostForm("username")
	email := c.PostForm("email")

	dbq := `
		INSERT INTO users (firebase_uid, username, email)
		VALUES (?, ?, ?)`

	result, err := models.DB.Exec(dbq, firebase_uid, username, email)
	if err != nil {
		panic(err)
	}

	user_id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"user_id": user_id})

}

func ListAllTasks(c *gin.Context) {
	dbq := `
		CALL get_all_tasks(?)`
	firebase_uid := c.Param("firebase_uid")

	rows, err := models.DB.Query(dbq, firebase_uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		panic(err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID, &task.UserID, &task.TaskName, &task.TaskDetail, &task.CreatedAt)
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetTask(c *gin.Context) {
	firebase_uid := c.Param("firebase_uid")
	task_id := c.Param("task_id")

	dbq := `
		CALL get_task(?, ?)`
	var task models.Task

	row := models.DB.QueryRow(dbq, firebase_uid, task_id)
	err := row.Scan(&task.TaskID, &task.UserID, &task.TaskName, &task.TaskDetail, &task.CreatedAt)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func CreateNewTask(c *gin.Context) {
	firebase_uid := c.Param("firebase_uid")
	task_name := c.PostForm("task_name")
	task_detail := c.PostForm("task_detail")

	var task_id int

	dbq := `
		CALL create_task(?, ?, ?)`

	err := models.DB.QueryRow(dbq, firebase_uid, task_name, task_detail).Scan(&task_id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"task_id": task_id})

}

func UpdateTask(c *gin.Context) {
	firebase_uid := c.Param("firebase_uid")
	task_id := c.Param("task_id")
	task_name := c.PostForm("task_name")
	task_detail := c.PostForm("task_detail")

	dbq := `
		CALL update_task(?, ?, ?, ?)`

	_, err := models.DB.Exec(dbq, firebase_uid, task_id, task_name, task_detail)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func DeleteTask(c *gin.Context) {
	firebase_uid := c.Param("firebase_uid")
	task_id := c.Param("task_id")

	dbq := `
		CALL delete_task(?, ?)`

	_, err := models.DB.Exec(dbq, firebase_uid, task_id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
