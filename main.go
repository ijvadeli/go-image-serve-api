package main

import (
	"github.com/gin-gonic/gin"
)

func handleImage(c *gin.Context) {
	// Check if parameter is equal to one of the files in img folder
	// Path parameters are used to identify something specific
	// Query parameters are used to sort or filter those resources
	image := c.Param("image")
	c.File("img/" + image)
}

func main() {
	r := gin.Default()

	r.GET("/img/:image", handleImage)

	r.Run(":8080")
}
