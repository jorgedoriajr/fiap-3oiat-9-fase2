package read

const tableIngredientTypeColumns = `
		  name,
		  optional,
		  max_qtd,
		  produtct_category
`
const FindIngredientTypeByName = `
		SELECT
			` + tableIngredientTypeColumns + `
		FROM ingredients_types
		WHERE name = $1
		LIMIT 1`
const FindIngredientTypeByProductCategory = `
		SELECT
			` + tableIngredientTypeColumns + `
		FROM ingredients_types
		WHERE produtct_category = $1
`

const FindIngredientTypeAll = `
		SELECT
			*
		FROM ingredients_types`
