package read

const tableColumns = `
		  id,
		  name,
		  amount,
		  description,
		  category,
		  menu,
		  ingredients,
		  created_at,
		  updated_at
`
const FindProductByID = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE id = $1
		LIMIT 1`

const FindProductByName = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE name = $1
		LIMIT 1`
