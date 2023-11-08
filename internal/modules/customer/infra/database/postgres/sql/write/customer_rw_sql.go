package write

const InsertCustomerRW = `
		INSERT INTO customer (
		  cpf, 
		  phone,            
		  name,             
		  email,            
		  opt_in_promotion,       
		  created_at,
		  updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
`
