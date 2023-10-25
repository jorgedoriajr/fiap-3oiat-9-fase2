package write

const InsertProductRW = `
		INSERT INTO customer (
		  id,
		  name,
		  amount,
		  description,
		  category,
		  menu,
		  ingredients,
		  created_at,
		  updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`
