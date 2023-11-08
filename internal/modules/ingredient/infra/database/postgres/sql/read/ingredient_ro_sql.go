package read

const tableColumns = `
		  id,
		  number,
		  name,
		  amount,
          type
`
const FindAllIngredients = `
		SELECT
			` + tableColumns + `
		FROM ingredient
		WHERE active = true
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
		AND active = true
`

const FindIngredientsByProductID = `
	SELECT
		i.id AS id,
		i.number as number,
		i.name AS name,
		i.type AS type,
		pi.amount as total_amount,
		pi.quantity as quantity
	FROM
		ingredient AS i
	JOIN
		product_ingredient AS pi ON i.id = pi.ingredient_id
	WHERE
		pi.product_id = $1;
`
