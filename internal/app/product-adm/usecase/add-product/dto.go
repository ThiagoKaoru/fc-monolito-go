package addproduct

import (
	"time"

	"github.com/google/uuid"
)

type AddProductInputDTO struct {
	Name          string
	Description   string
	PurchasePrice uint
	Stock         int
}

type AddProductOutputDTO struct {
	Id            uuid.UUID
	Name          string
	Description   string
	PurchasePrice uint
	Stock         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
