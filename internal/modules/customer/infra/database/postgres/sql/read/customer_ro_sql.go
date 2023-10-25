package read

const tableColumns = `
			cpf,              
			phone,            
			name,             
			email,            
			opt_in_promotion,       
			created_at,
			updated_at
`
const FindCustomerByCpf = `
		SELECT
			` + tableColumns + `
		FROM customer
		WHERE cpf = $1
		LIMIT 1`
