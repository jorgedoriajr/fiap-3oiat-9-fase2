package read

const categoryTableColumns = `
		  id,
		  name,
`
const FindAllProductCategories = `
		SELECT
			` + categoryTableColumns + `
		FROM product_category
		`
const FindProductCategoryByID = `
		SELECT
			` + categoryTableColumns + `
		FROM product_category
		WHERE id = $1
		LIMIT 1`

const FindProductCategoryByName = `
		SELECT
			` + categoryTableColumns + `
		FROM product_category
		WHERE name = $1
		`
