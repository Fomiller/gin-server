package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	router := gin.Default()
	// router.LoadHTMLFiles("templates/index.gohtml")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", indexHandler)
	router.GET("/welcome", welcomeHandler)
	router.GET("/welcome2", welcomeHandler2)
	router.GET("/file", fileHandler)

	// secret := router.Group("/secret")
	// {
	// 	secret.GET("/clubhouse", secretHandler)
	// 	secret.GET("/clubhouse/:password", clubHouseHandler)
	// }
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorized.GET("/secrets", secretsHandler)

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
	data := gin.H{
		"firstName": "Forrest",
		"lastName":  "Miller",
	}

	c.HTML(http.StatusOK, "index.gohtml", data)
}

// func secretHandler(c *gin.Context) {
// 	c.HTML(http.StatusOK, "secret.gohtml", nil)
// }

// func clubHouseHandler(c *gin.Context) {
// 	password := c.Query("password")
// 	c.HTML(http.StatusOK, "clubHouse.gohtml", password)
// }

func secretsHandler(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.HTML(http.StatusOK, "clubhouse.gohtml", gin.H{
			"user":   user,
			"secret": secret,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user":   user,
			"secret": "NO SECRET :(",
		})
	}

}
