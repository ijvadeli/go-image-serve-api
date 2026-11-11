package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Simple home handler displaying client ip
func homeHandler(c *gin.Context) {
	clientIP := c.ClientIP()

	c.HTML(http.StatusOK, "templates/index.tmpl", gin.H{
		"clientip": clientIP,
	})
}

func handleImage(c *gin.Context) {
	// Create image parameter
	image := c.Param("image")

	// Serve html file based off of index.tmpl template and using the image parameter
	c.HTML(http.StatusOK, "templates/image.tmpl", gin.H{
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
	// Routing
	r.GET("/", homeHandler)
	r.GET("/:image", handleImage)

	r.Run(":8080")
}
