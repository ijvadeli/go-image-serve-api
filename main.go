package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleImage(c *gin.Context) {
	// Check if parameter is equal to one of the files in img folder
	// Path parameters are used to identify something specific
	// Query parameters are used to sort or filter those resources
	image := c.Param("image")

	// Serve html file based off of index.tmpl template
	c.HTML(http.StatusOK, "templates/index.tmpl", gin.H{
		"title": image,
		"image": "/assets/" + image,
	})
}

func main() {
	r := gin.Default()

	// Set trusted proxies using default gateway/admin address
	r.SetTrustedProxies([]string{"192.168.1.2"})

	// Load html files identified by global pattern, I use tmpl since gin recommends it
	// tmpl is template file used alot by go
	r.LoadHTMLGlob("templates/*")

	// Mapping between URL path and filesystem path to display images
	r.Static("/assets", "./assets")

	r.GET("/:image", handleImage)

	r.Run(":8080")
}
