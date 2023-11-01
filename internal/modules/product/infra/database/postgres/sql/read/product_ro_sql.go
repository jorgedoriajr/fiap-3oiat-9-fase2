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
		WHERE active = true
`

const FindProductByID = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE id = $1
		LIMIT 1`

const FindProductByNumber = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE number = $1
		LIMIT 1`

const FindProductByCategory = `
		SELECT
			` + tableColumns + `
		FROM product
		WHERE category = $1
		AND active = true
`

const FindProductByOrderID = `
	SELECT
		p.id,
		p.name,
		p.number,
		p.description,
		p.category,
		p.img_path,
		op.amount as total_amount,
		op.quantity
	FROM
	    order_product AS op
	JOIN
		 product AS p ON p.id = op.product_id
	WHERE
		op.order_id = $1;
`
