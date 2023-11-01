package write

const InsertProductRW = `
		INSERT INTO product (
		  id,
		  name,
		  amount,
		  description,
		  category,
		  menu,
		  img_path,
		  created_at,
		  updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
`

const UpdateProductRW = `
	UPDATE product
	SET
	  name = COALESCE($2, name),
	  description = COALESCE($3, description),
	  category = COALESCE($4, category),
	  menu = COALESCE($5, menu),
	  img_path = COALESCE($6, img_path)
	WHERE number = $1
`

const InsertProductIngredientRW = `
		INSERT INTO product_ingredient (
		  id, 
		  product_id, 
		  ingredient_id, 
		  quantity, 
		  amount
		) VALUES ($1, $2, $3, $4, $5) RETURNING id
`

const InactiveProductById = `
		UPDATE product
			SET active = false
			WHERE id = $1
`
