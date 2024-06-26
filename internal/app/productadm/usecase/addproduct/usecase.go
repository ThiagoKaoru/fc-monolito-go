package productusecase

import (
	"github.com/google/uuid"
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/domain"
	productadm_repository "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/repository"
	baseEntity "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/entity"
	IdValueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/valueobject"
)

type AddProduct struct {
	productRepository *productadm_repository.ProductRepository
}

func NewAddProductUseCase(productRepository *productadm_repository.ProductRepository) *AddProduct {
	return &AddProduct{
		productRepository: productRepository,
	}
}
func (ap *AddProduct) Execute(input AddProductInputDTO) error {
	product := productDomain.NewProduct(productDomain.Product{
		BaseEntity:    baseEntity.BaseEntity{ID: IdValueObject.ID{ID: uuid.New()}},
		Name:          input.Name,
		Description:   input.Description,
		PurchasePrice: input.PurchasePrice,
		Stock:         input.Stock,
	})
	productRepository := ap.productRepository.Add(product)

	return productRepository
}
