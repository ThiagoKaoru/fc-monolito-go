package productadm_facade

import (
	"context"

	addproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/usecase/addproduct"
	findproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/usecase/findproduct"
)

type Facade interface {
	addProduct(context.Context, addproduct.AddProductInputDTO) error
	checkStock(context.Context, findproduct.FindProductInputDTO) (findproduct.FindProductOutputDTO, error)
}
