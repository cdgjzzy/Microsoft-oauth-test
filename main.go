package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/redirect", redirect)
	r.GET("/code", code)
	r.GET("/display", display)

	r.Run(":5001")
}
