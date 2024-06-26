package productfactory

import (
	productadm_facade "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/facade"
	productadm_repository "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/repository"
	productadm_usecase_add "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/usecase/addproduct"
	productadm_usecase_find "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/usecase/findproduct"
	"github.com/thiagokaoru/fc-monolito-go/internal/pkg/usecase"
	"gorm.io/gorm"
)

func ProductAdmFacadeFactory(db *gorm.DB) productadm_facade.ProductAdmFacade {
	productRepository := productadm_repository.NewProductRepository(db)
	productAddUseCase := productadm_usecase_add.NewAddProductUseCase(productRepository)
	productFindUseCase := productadm_usecase_find.NewFindProductUseCase(productRepository)

	productFacade := productadm_facade.NewProductAdmFacade(
		usecase.UseCaseNoOutputInterface[productadm_usecase_add.AddProductInputDTO](productAddUseCase),
		usecase.UseCaseInterface[productadm_usecase_find.FindProductInputDTO, productadm_usecase_find.FindProductOutputDTO](productFindUseCase),
	)
	return *productFacade
}
