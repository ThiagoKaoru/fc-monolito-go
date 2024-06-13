package facade

import (
	"context"
)

type Facade interface {
	addProduct(context.Context, AddProductFacadeInputDTO) error
	checkStock(context.Context, CheckStockFacadeInputDTO) (CheckStockFacadeOutputDTO, error)
}

type AddProductFacadeInputDTO struct {
	ID            *string
	Name          string
	Description   string
	PurchasePrice uint16
	Stock         uint16
}

type CheckStockFacadeInputDTO struct {
	ProductId string
}

type CheckStockFacadeOutputDTO struct {
	ProductId string
	Stock     uint16
}
