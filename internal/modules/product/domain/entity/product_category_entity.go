package entity

import "github.com/google/uuid"

type ProductCategoryEntity struct {
	ID   uuid.UUID
	Name string
}
