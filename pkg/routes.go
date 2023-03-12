package pkg

import (
	"github.com/gin-gonic/gin"
)

//defining routes
func Routes() {
	router := gin.Default()
	router.GET("/products", listProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/productsCreate", createProduct)
	router.DELETE("/productsDelete/:id", deleteProduct)
	router.GET("/search", searchProducts)
	router.GET("/sort", sortProducts)
	router.PUT("/productsUpdate/:id", updateProduct)


	// Start the server
	router.Run(":8080")
}