package read

const categoryTableColumns = `
		  name,
		  accept_custom
`
const FindAllProductCategories = `
		SELECT
			` + categoryTableColumns + `
		FROM product_category
		`

const FindProductCategoryByName = `
		SELECT
			` + categoryTableColumns + `
		FROM product_category
		WHERE name = $1
		`
