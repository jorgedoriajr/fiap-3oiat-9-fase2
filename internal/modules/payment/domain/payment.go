package domain

import "github.com/google/uuid"

type Payment struct {
	Id      uuid.UUID
	OrderId uuid.UUID
	Data    string
}
