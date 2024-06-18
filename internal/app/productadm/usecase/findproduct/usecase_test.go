package productadm_usecase_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/domain"
	productadm_repository "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/repository"
	productadm_usecase "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/usecase/findproduct"
	IdValueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/value-object"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FindProductSuite struct {
	suite.Suite
	db *gorm.DB
}

func (s *FindProductSuite) SetupTest() {
	s.db, _ = gorm.Open(sqlite.Open(":memory"), &gorm.Config{})
	s.NoError(s.db.AutoMigrate(&productadm_usecase.FindProductOutputDTO{}))
}

func (s *FindProductSuite) TestFindProduct() {
	product := productDomain.Product{
		Name:          "Product 1",
		Description:   "Product 1 description",
		PurchasePrice: 100,
		Stock:         10,
	}
	uniqueID := uuid.New()
	product.ID = IdValueObject.NewID(uniqueID)

	s.db.Create(&product)
	repositoryProduct := productadm_repository.NewProductRepository(s.db)
	result, err := productadm_usecase.NewFindProductUseCase(repositoryProduct).Execute(productadm_usecase.FindProductInputDTO{
		ProductID: product.ID,
	})

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), product.Name, result.Name)
	assert.Equal(s.T(), product.Description, result.Description)
	assert.Equal(s.T(), product.PurchasePrice, result.PurchasePrice)
	assert.Equal(s.T(), product.Stock, result.Stock)
}
