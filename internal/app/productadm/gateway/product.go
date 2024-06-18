package productGateway

import (
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/productadm/domain"
	IdValueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/value-object"
)

type ProductGateway interface {
	Add(product *productDomain.Product) error
	Find(id IdValueObject.ID) (*productDomain.Product, error)
}
