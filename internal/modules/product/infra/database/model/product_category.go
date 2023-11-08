package model

import "hamburgueria/internal/modules/product/domain"

type ProductCategory struct {
	Name         string
	AcceptCustom bool
}

func (p ProductCategory) ToDomain() *domain.ProductCategory {
	return &domain.ProductCategory{
		Name:         p.Name,
		AcceptCustom: p.AcceptCustom,
	}
}
