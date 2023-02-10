package main

import (
	"gin_web_test/src/router"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	r = gin.Default()
	router.Init(r)
	// Listen and Server in 0.0.0.0:8080
	r.Use(gin.Logger())
	r.Run(":8080")
}
