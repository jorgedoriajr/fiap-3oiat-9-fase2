package presenter

import (
	"hamburgueria/internal/modules/product/usecase/result"
	"hamburgueria/internal/web/api/rest/v1/productcategory/response"
)

func ProductCategoriesResponseFromResult(result []result.FindProductCategoryResult) []response.ProductCategoryResponse {
	var productCategoryResponse []response.ProductCategoryResponse
	for _, categoryResult := range result {
		productCategoryResponse = append(productCategoryResponse,
			response.ProductCategoryResponse{
				Name:         categoryResult.Name,
				AcceptCustom: categoryResult.AcceptCustom,
			},
		)
	}
	return productCategoryResponse
}
