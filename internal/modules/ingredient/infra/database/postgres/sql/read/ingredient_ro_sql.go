package read

const tableColumns = `
		  id,
		  name,
		  amount,
          type
`
const FindAllIngredients = `
		SELECT
			` + tableColumns + `
		FROM ingredient
		`
const FindIngredientByID = `
		SELECT
			` + tableColumns + `
		FROM ingredient
		WHERE id = $1
		LIMIT 1`

const FindIngredientsByType = `
		SELECT
			` + tableColumns + `
		FROM ingredient
		WHERE type = $1
		`
