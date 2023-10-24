package read

const FindCustomerByCpf = `
		SELECT
			cpf,              
			phone,            
			name,             
			email,            
			opt_in_promotion,       
			created_at,
			updated_at
		FROM customer
		WHERE cpf = $1
		LIMIT 1`
