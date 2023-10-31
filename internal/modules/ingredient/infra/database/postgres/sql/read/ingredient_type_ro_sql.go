package read

const tableIngredientTypeColumns = `
		  name
`
const FindIngredientTypeByName = `
		SELECT
			` + tableIngredientTypeColumns + `
		FROM ingredient_type
		WHERE name = $1
		LIMIT 1`

const FindIngredientTypeByProductCategory = `
		SELECT
			` + tableIngredientTypeColumns + `
		FROM ingredient_type it
		INNER JOIN ingredient_type_product_category itpc on it.name = itpc.ingredient_type 
		WHERE product_category = $1
`

const FindIngredientTypeAll = `
		SELECT
			name
		FROM ingredient_type`
