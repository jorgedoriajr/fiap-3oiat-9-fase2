package read

const tableColumns = `
		  id,
		  customer_id, 
		  amount,      
		  status,      
		  created_at,
		  updated_at
`
const FindAllOrders = `
		SELECT
			` + tableColumns + `
		FROM "order"
		ORDER BY created_at
		`
const FindOrderById = `
		SELECT
			` + tableColumns + `
		FROM "order"
		WHERE id = $1
		LIMIT 1`

const FindOrderByStatus = `
		SELECT
			` + tableColumns + `
		FROM "order"
		WHERE status = $1
		ORDER BY created_at`
