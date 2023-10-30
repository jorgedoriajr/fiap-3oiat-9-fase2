package write

const InsertProductRW = `
		INSERT INTO product (
		  name,
		  amount,
		  description,
		  category,
		  menu,
		  created_at,
		  updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
`
