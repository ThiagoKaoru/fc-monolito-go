package productadm_facade

import (
	"context"

	addproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/usecase/add-product"
	findproduct "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/usecase/find-product"
)

type Facade interface {
	addProduct(context.Context, addproduct.AddProductInputDTO) error
	checkStock(context.Context, findproduct.FindProductInputDTO) (findproduct.FindProductOutputDTO, error)
}
