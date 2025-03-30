package main

import (
	"task_manager_api_auth/data"
	"task_manager_api_auth/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// setup mongoDB
	data.IntializeMongoBD()

	// run and get the router setup router function
	r := gin.Default()

	router.SetUpRouter(r)
	r.Run(":8080")

}
