package productadm_usecase

import IdValueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/valueobject"

type FindProductInputDTO struct {
	ProductID IdValueObject.ID
}
type FindProductOutputDTO struct {
	ProductID     IdValueObject.ID
	Name          string
	Description   string
	PurchasePrice uint
	Stock         int
}
