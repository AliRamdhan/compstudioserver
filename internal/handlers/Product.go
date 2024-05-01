package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/AliRamdhan/compstudioserver/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductServices
}

func NewProductHandler(ps *services.ProductServices) *ProductHandler {
	return &ProductHandler{productService: ps}
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product

	// Parse form fields
	name := c.PostForm("ProductName")
	link := c.PostForm("ProductLink")
	priceStr := c.PostForm("ProductPrice")

	// Parse the price string to uint
	price, err := strconv.ParseUint(priceStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
		return
	}

	// Get file header from context
	fileHeaderInterface, exists := c.Get("fileHeader")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File header not found in context"})
		return
	}
	fileHeader := fileHeaderInterface.(*multipart.FileHeader)

	// Bind form fields to the product struct
	product.ProductName = name
	product.ProductLink = link
	product.ProductPrice = uint(price)

	// Call service method to create product
	err = ph.productService.CreateProduct(&product, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

func (ph *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := ph.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All Products", "Product": products})
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productIDStr := c.Param("productId")
	// productID, err := uuid.Parse(productIDStr)
	var productID uint
	_, err := fmt.Sscanf(productIDStr, "%d", &productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := ph.productService.UpdateProduct(productID, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": product})
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	productIDStr := c.Param("productId")
	// productID, err := uuid.Parse(productIDStr)
	var productID uint
	_, err := fmt.Sscanf(productIDStr, "%d", &productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := ph.productService.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
