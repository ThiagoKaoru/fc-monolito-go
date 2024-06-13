package IdValueObject

import "github.com/google/uuid"

type ID struct {
	ID uuid.UUID `gorm:"primarykey"`
}

func NewID(Id uuid.UUID) ID {
	return ID{ID: Id}
}
func (ov *ID) GetID() uuid.UUID {
	return ov.ID
}
