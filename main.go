package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./dist/assets")
	router.LoadHTMLGlob("./dist/index.html")
	router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	//show homepage
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Run(":8080")
}
