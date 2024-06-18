package facade_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/factory"
	"github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/repository"
	addproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/usecase/add-product"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductRepositorySuite struct {
	suite.Suite
	db *gorm.DB
}

func (s *ProductRepositorySuite) SetupTest() {
	s.db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.NoError(s.db.AutoMigrate(&repository.Product{}))
}

func (s *ProductRepositorySuite) TestInsertAndVerifyProduct() {
	product := addproduct.AddProductInputDTO{
		Name:          "Product 1",
		Description:   "Product 1 description",
		PurchasePrice: 100,
		Stock:         10,
	}
	productFacade := factory.ProductAdmFacadeFactory(s.db)
	result := productFacade.AddProduct(product)
	assert.NoError(s.T(), result)

	var foundProduct addproduct.AddProductOutputDTO
	result = s.db.First(&foundProduct, "Name =?", product.Name).Error
	assert.NoError(s.T(), result)
	assert.Equal(s.T(), product.Name, foundProduct.Name)
	assert.Equal(s.T(), product.Description, foundProduct.Description)
	assert.Equal(s.T(), product.PurchasePrice, foundProduct.PurchasePrice)
	assert.Equal(s.T(), product.Stock, foundProduct.Stock)
}
