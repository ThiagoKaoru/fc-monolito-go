package factory

import (
	"github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/facade"
	"github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/repository"
	addproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/usecase/add-product"
	findproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/usecase/find-product"
	"github.com/thiagokaoru/fc-monolito-go/internal/pkg/usecase"
	"gorm.io/gorm"
)

func ProductAdmFacadeFactory(db *gorm.DB) facade.ProductAdmFacade {
	productRepository := repository.NewProductRepository(db)
	productAddUseCase := addproduct.NewAddProductUseCase(productRepository)
	productFindUseCase := findproduct.NewFindProductUseCase(productRepository)

	productFacade := facade.NewProductAdmFacade(
		usecase.UseCaseNoOutputInterface[addproduct.AddProductInputDTO](productAddUseCase),
		usecase.UseCaseInterface[findproduct.FindProductInputDTO, findproduct.FindProductOutputDTO](productFindUseCase),
	)
	return *productFacade
}
