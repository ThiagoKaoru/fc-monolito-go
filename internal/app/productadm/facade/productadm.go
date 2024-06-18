package productfacade

import (
	addproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/usecase/addproduct"
	findproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/usecase/findproduct"
	"github.com/thiagokaoru/fc-monolito-go/internal/pkg/usecase"
)

type ProductAdmFacade struct {
	addUseCase       usecase.UseCaseNoOutputInterface[addproduct.AddProductInputDTO]
	findStockUseCase usecase.UseCaseInterface[findproduct.FindProductInputDTO, findproduct.FindProductOutputDTO]
}

func NewProductAdmFacade(addUseCase usecase.UseCaseNoOutputInterface[addproduct.AddProductInputDTO], findStockUseCase usecase.UseCaseInterface[findproduct.FindProductInputDTO, findproduct.FindProductOutputDTO]) *ProductAdmFacade {
	return &ProductAdmFacade{
		addUseCase:       addUseCase,
		findStockUseCase: findStockUseCase,
	}
}

func (f *ProductAdmFacade) AddProduct(input addproduct.AddProductInputDTO) error {
	err := f.addUseCase.Execute(input)
	return err
}

func (f *ProductAdmFacade) FindStock(input findproduct.FindProductInputDTO) (*findproduct.FindProductOutputDTO, error) {
	var output findproduct.FindProductOutputDTO
	output, err := f.findStockUseCase.Execute(input)
	return &output, err
}
