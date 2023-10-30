package write

const InsertProductCategoryRW = `
		INSERT INTO product (
		  id,
		  name,
		) VALUES ($1, $2)
`
