package productadm_repository

import (
	"errors"

	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/domain"
	IdValueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/valueobject"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Add(product *productDomain.Product) error {
	result := r.db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *ProductRepository) Find(id IdValueObject.ID) (*productDomain.Product, error) {
	var product productDomain.Product
	result := r.db.First(&product, "id =?", id.ID.String())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Product with id " + id.ID.String() + " not found")
		}
		return nil, result.Error
	}
	return &product, nil
}
