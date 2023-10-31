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
		WHERE id = $1
		LIMIT 1`

const FindIngredientsByProductID = `
	SELECT
		i.id AS id,
		i.name AS name,
		i.amount AS amount,
		i.type AS type,
		pi.quantity as quantity
	FROM
		ingredient AS i
	JOIN
		product_ingredient AS pi ON i.id = pi.ingredient_id
	WHERE
		pi.product_id = $1;
`
