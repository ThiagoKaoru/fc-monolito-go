package facade

import (
	"context"

	"github.com/thiagokaoru/fc-monolito-go/internal/pkg/usecase"
)

type ProductAdmFacade struct {
	addUseCase        usecase.UseCaseNoOutputInterface[AddProductFacadeInputDTO]
	checkStockUseCase usecase.UseCaseInterface[CheckStockFacadeInputDTO, CheckStockFacadeOutputDTO]
}

func NewProductAdmFacade(addUseCase usecase.UseCaseNoOutputInterface[AddProductFacadeInputDTO], checkStockUseCase usecase.UseCaseInterface[CheckStockFacadeInputDTO, CheckStockFacadeOutputDTO]) *ProductAdmFacade {
	return &ProductAdmFacade{
		addUseCase:        addUseCase,
		checkStockUseCase: checkStockUseCase,
	}
}

func (f *ProductAdmFacade) AddProduct(ctx context.Context, input AddProductFacadeInputDTO) error {
	err := f.addUseCase.Execute(ctx, input)
	return err
}

func (f *ProductAdmFacade) CheckStock(ctx context.Context, input CheckStockFacadeInputDTO) (*CheckStockFacadeOutputDTO, error) {
	var output CheckStockFacadeOutputDTO
	output, err := f.checkStockUseCase.Execute(ctx, input)
	return &output, err
}
