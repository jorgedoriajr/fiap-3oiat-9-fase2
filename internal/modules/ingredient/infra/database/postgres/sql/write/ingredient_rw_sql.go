package write

const InsertIngredientRW = `
	INSERT INTO ingredient (
	  id,
	  name,
	  amount,
	  type
	) VALUES ($1, $2, $3, $4)
	ON CONFLICT (name) DO NOTHING;
`
