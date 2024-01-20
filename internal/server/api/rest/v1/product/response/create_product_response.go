package response

import (
	"github.com/google/uuid"
)

type ProductCreatedResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Menu        bool      `json:"menu"`
	ImgPath     string    `json:"imgPath"`
}
