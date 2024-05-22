package valueObject

import "github.com/google/uuid"

type ID struct {
	Id uuid.UUID
}

func NewID(Id uuid.UUID) *uuid.UUID {
	return &Id
}
func (ov *ID) GetID() uuid.UUID {
	return ov.Id
}
