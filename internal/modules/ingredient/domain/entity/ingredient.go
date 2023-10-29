package entity

import (
	"github.com/google/uuid"
)

type Ingredient struct {
	Id     uuid.UUID
	Name   string
	Amount string
	Type   string
}
