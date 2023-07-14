package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Features
	// Add todo
	// Store created_at for each todo
	// Store completed_at for each todo
	// Get todo by id
	// List todo
	// Mark todo as done
	// Delete todo

	// API Call Add example
	// curl -X POST -H "Content-Type: application/json" -d '{"title":"Buy milk"}' http://localhost:8080/todos
	// Response Code 201 Created
	// Response Body {"id":"6807E49C-D68F-4D38-A990-6FACCB0E51A3","title":"Buy milk","created_at":"2020-12-20T15:04:05Z","completed_at": null}

	// API Call Get by id example
	// curl -X GET http://localhost:8080/todos/6807E49C-D68F-4D38-A990-6FACCB0E51A3
	// Response Code 200 OK
	// Response Body {"id":"6807E49C-D68F-4D38-A990-6FACCB0E51A3","title":"Buy milk","created_at":"2020-12-20T15:04:05Z","completed_at": null}

	// API Call List example
	// curl -X GET http://localhost:8080/todos
	// Response Code 200 OK
	// Response Body [{"id":"6807E49C-D68F-4D38-A990-6FACCB0E51A3","title":"Buy milk","created_at":"2020-12-20T15:04:05Z","completed_at": null}]

	// API Call Mark as done example
	// curl -X PUT http://localhost:8080/todos/6807E49C-D68F-4D38-A990-6FACCB0E51A3/done
	// Response Code 200 OK
	// Response Body {"id":"6807E49C-D68F-4D38-A990-6FACCB0E51A3","title":"Buy milk","created_at":"2020-12-20T15:04:05Z","completed_at": "2020-12-20T15:04:05Z"}

	// API Call Delete example
	// curl -X DELETE http://localhost:8080/todos/6807E49C-D68F-4D38-A990-6FACCB0E51A3
	// Response Code 200 OK

	router := gin.Default()
	router.POST("/todos", func(c *gin.Context) {
	})
	router.GET("/todos/:id", func(c *gin.Context) {
		// c.Params("id")
	})
	router.GET("/todos", func(c *gin.Context) {
	})
	router.PUT("/todos/:id/done", func(c *gin.Context) {
	})
	router.DELETE("/todos/:id", func(c *gin.Context) {
	})

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
