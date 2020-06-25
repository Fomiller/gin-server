package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.gohtml")
	router.GET("/", indexHandler)
	router.GET("/welcome", welcomeHandler)
	router.GET("/welcome2", welcomeHandler2)
	router.GET("/file", fileHandler)

	router.Run()
}

func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HELLO WORLD",
		"number":  123123123123,
	})
}

func welcomeHandler(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func welcomeHandler2(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.JSON(http.StatusOK, gin.H{
		"firstName": firstname,
		"lastName":  lastname,
	})
}

func fileHandler(c *gin.Context) {
	f
}
