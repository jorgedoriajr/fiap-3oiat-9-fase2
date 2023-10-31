package response

import (
	"github.com/google/uuid"
)

type ProductCreatedResponse struct {
	Id          uuid.UUID
	Name        string
	Amount      int
	Description string
	Category    string
	Menu        bool
	ImgPath     string
}
