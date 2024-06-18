package productusecase

import (
	productadm_repository "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/repository"
)

type FindProduct struct {
	productRepository *productadm_repository.ProductRepository
}

func NewFindProductUseCase(productRepository *productadm_repository.ProductRepository) *FindProduct {
	return &FindProduct{
		productRepository: productRepository,
	}
}
func (ap *FindProduct) Execute(input FindProductInputDTO) (FindProductOutputDTO, error) {
	productRepository, err := ap.productRepository.Find(input.ProductID)
	if err != nil {
		return FindProductOutputDTO{}, err
	}
	return FindProductOutputDTO{
		ProductID:     productRepository.ID,
		Name:          productRepository.Name,
		Description:   productRepository.Description,
		PurchasePrice: productRepository.PurchasePrice,
		Stock:         productRepository.Stock,
	}, nil
}
