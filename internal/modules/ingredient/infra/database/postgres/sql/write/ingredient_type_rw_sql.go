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
const UpdateIngredientTypwRW = `
	UPDATE ` + tableNameIngredientType + `
	SET name = $1,
		optional = $2,
		max_qtd = $3,
		product_category = $5
	WHERE name = $1
`
