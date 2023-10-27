package write

const tableNameIngredientType string = "ingredient_types"

const InsertIngredientTypeRW = `
		INSERT INTO ` + tableNameIngredientType + ` (
		name,
		optional,
		max_qtd,
		produtct_category,
		) VALUES ($1, $2, $3, $4)
`
