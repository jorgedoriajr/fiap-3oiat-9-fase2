package write

const InsertOrderRW = `
		INSERT INTO "order" (
			id,
			customer_id, 
			amount,            
			status,                    
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
`

const InsertOrderProductRW = `
		INSERT INTO order_product (
			id,
			order_id,
		    product_id,            
			quantity,  
		    amount
		)
		VALUES ($1, $2, $3, $4, $5)
`

const InsertOrderHistoryRW = `
		INSERT INTO order_history (
			id,
		    order_id,            
			status,                    
			created_at,
			change_by
		)
		VALUES ($1, $2, $3, $4, $5)
`

const UpdateOrderPayment = `
		UPDATE "order"
		SET 
		    payment_id = $1,
		    status = $2
		WHERE id = $3
`
