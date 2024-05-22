package entity

import (
	"time"

	valueObject "github.com/thiagokaoru/fc-monolito-go/internal/pkg/domain/value-object"
)

type AggregateRoot interface {
}
type BaseEntity struct {
	ID        valueObject.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBaseEntity(id *valueObject.ID, createdAt, updatedAt time.Time) *BaseEntity {
	return &BaseEntity{
		ID:        *id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (b *BaseEntity) GetID() valueObject.ID {
	return b.ID
}

func (b *BaseEntity) GetCreatedAt() time.Time {
	return b.CreatedAt
}

func (b *BaseEntity) GetUpdatedAt() time.Time {
	return b.UpdatedAt
}

func (b *BaseEntity) SetUpdatedAt(updatedAt time.Time) {
	b.UpdatedAt = updatedAt
}
