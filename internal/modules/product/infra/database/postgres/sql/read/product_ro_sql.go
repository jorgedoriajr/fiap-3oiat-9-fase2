package read

const tableColumns = `
		  id,
		  name,
		  number,
		  amount,
		  description,
		  category,
		  menu,
		  img_path,	
		  created_at,
		  updated_at
`
const FindAllProducts = `
		SELECT
			` + tableColumns + `
		FROM product
		`
const FindProductByID = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE id = $1
		LIMIT 1`

const FindProductByIDWithIngredients = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE id = $1
		LIMIT 1`

const FindProductByCategory = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE category = $1
		`
