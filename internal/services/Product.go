package services

import (
	"mime/multipart"
	"time"

	"github.com/AliRamdhan/compstudioserver/config"
	"github.com/AliRamdhan/compstudioserver/internal/model"
	"github.com/AliRamdhan/compstudioserver/utils"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type ProductServices struct{}

func NewProductServices() *ProductServices {
	return &ProductServices{}
}

func (ps *ProductServices) CreateProduct(product *model.Product, fileHeader *multipart.FileHeader) error {
	// Validate product
	if err := validate.Struct(product); err != nil {
		return err
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Upload file to Cloudinary
	imageURL, err := utils.UploadToCloudinary(file, fileHeader.Filename)
	if err != nil {
		return err
	}

	// Assign Cloudinary URL to ProductImage field
	product.ProductImage = imageURL

	// Set creation timestamp
	product.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Create product in the database
	return config.DB.Create(product).Error
}

func (ps *ProductServices) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductServices) UpdateProduct(productID uint, updatedProduct *model.Product) error {
	// Find the product with the given ID
	var existingProduct model.Product
	if err := config.DB.First(&existingProduct, "product_id = ?", productID).Error; err != nil {
		return err // Product not found
	}

	// Update fields of existing product with the new values
	existingProduct.ProductName = updatedProduct.ProductName
	existingProduct.ProductPrice = updatedProduct.ProductPrice
	existingProduct.ProductLink = updatedProduct.ProductLink
	existingProduct.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Save the updated product
	return config.DB.Save(&existingProduct).Error
}

func (ps *ProductServices) DeleteProduct(productID uint) error {
	// Find the product with the given ID
	var product model.Product
	if err := config.DB.First(&product, "product_id = ?", productID).Error; err != nil {
		return err // Product not found
	}
	// Delete the product
	return config.DB.Delete(&product).Error
}
