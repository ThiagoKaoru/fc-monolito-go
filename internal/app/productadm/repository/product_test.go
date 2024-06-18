package productadm_repository

import (
	"testing"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/domain"
	IdValueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/value-object"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductRepositorySuite struct {
	suite.Suite
	db *gorm.DB
}

func (s *ProductRepositorySuite) SetupTest() {
	s.db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.NoError(s.db.AutoMigrate(&Product{}))
}

func (s *ProductRepositorySuite) TestInsertAndVerifyProduct() {
	product := productDomain.Product{
		Name:          "Product 1",
		Description:   "Product 1 description",
		PurchasePrice: 100,
		Stock:         10,
	}
	uniqueID := uuid.New()
	product.ID = IdValueObject.NewID(uniqueID)

	result := NewProductRepository(s.db).Add(&product)
	assert.NoError(s.T(), result)

	var foundProduct Product
	result = s.db.First(&foundProduct, "id =?", uniqueID).Error
	assert.NoError(s.T(), result)
	assert.Equal(s.T(), product.Name, foundProduct.Name)
	assert.Equal(s.T(), product.Description, foundProduct.Description)
	assert.Equal(s.T(), product.PurchasePrice, foundProduct.PurchasePrice)
	assert.Equal(s.T(), product.Stock, foundProduct.Stock)
}

func (s *ProductRepositorySuite) TestFindAProduct() {
	product := productDomain.Product{
		Name:          "Product 1",
		Description:   "Product 1 description",
		PurchasePrice: 100,
		Stock:         10,
	}
	uniqueID := uuid.New()
	product.ID = IdValueObject.NewID(uniqueID)
	err := s.db.Create(&product).Error
	assert.NoError(s.T(), err)

	result, err := NewProductRepository(s.db).Find(IdValueObject.ID{ID: uniqueID})
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), product.Name, result.Name)
	assert.Equal(s.T(), product.Description, result.Description)
	assert.Equal(s.T(), product.PurchasePrice, result.PurchasePrice)
	assert.Equal(s.T(), product.Stock, result.Stock)
}
func TestProductRepositorySuite(t *testing.T) {
	suite.Run(t, new(ProductRepositorySuite))
}
