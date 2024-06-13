package addproduct

import (
	"github.com/google/uuid"
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/domain"
	productGateway "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/gateway"
	baseEntity "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/entity"
	IdValueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/value-object"
)

type addProduct struct {
	productRepository *productGateway.ProductGateway
}

func NewAddProductUseCase(productRepository *productGateway.ProductGateway) *addProduct {
	return &addProduct{
		productRepository: productRepository,
	}
}
func (ap *addProduct) Execute(input AddProductInputDTO) AddProductOutputDTO {
	product := productDomain.NewProduct(productDomain.Product{
		BaseEntity:    baseEntity.BaseEntity{ID: IdValueObject.ID{ID: uuid.New()}},
		Name:          input.Name,
		Description:   input.Description,
		PurchasePrice: input.PurchasePrice,
		Stock:         input.Stock,
	})
	Add(*product)
	return AddProductOutputDTO{
		Id:            product.BaseEntity.ID.ID,
		Name:          product.Name,
		Description:   product.Description,
		PurchasePrice: product.PurchasePrice,
		Stock:         product.Stock,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}
}

func Add(input productDomain.Product) error {
	return nil
}
