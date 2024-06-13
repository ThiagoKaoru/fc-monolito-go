package productGateway

import (
	"github.com/google/uuid"
	productDomain "github.com/thiagokaoru/fc-monolito-go/internal/app/product-adm/domain"
)

type ProductGateway interface {
	Add(product productDomain.Product) error
	Find(id uuid.UUID) (chan productDomain.Product, error)
}
