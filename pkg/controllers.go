package pkg

import (
	"net/http"
	"strconv"

	md "github.com/LeilaBeken/golang_ass_3/models"
	"github.com/gin-gonic/gin"
)

func listProducts(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	products := []md.Book{}
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	productID := c.Param("id")
	product := md.Book{}
	db.First(&product, productID)
	if product.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	var productData md.Book
	if err := c.ShouldBindJSON(&productData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&productData)
	c.JSON(http.StatusOK, productData)
}

func deleteProduct(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	productID := c.Param("id")
	existingProduct := md.Book{}
	db.First(&existingProduct, productID)
	if existingProduct.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	db.Delete(&existingProduct)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func updateProduct(c *gin.Context) {
	var product md.Book
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	productID := c.Param("id")

	// Bind the JSON request body to the product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the product in the database by ID
	if err := db.Where("id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Convert the price string to an integer
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}

	// Update the product with the new information
	product.Title = c.PostForm("title")
	product.Description = c.PostForm("author")
	product.Price = price

	// Save the changes to the database
	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, product)
}
