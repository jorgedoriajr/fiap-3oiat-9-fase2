CREATE TABLE IF NOT EXISTS "payment" (
	id       uuid.UUID
	order_id uuid.UUID
	Data     varchar(255)
);

CREATE TABLE IF NOT EXISTS "payment_integration_log" (
	id                     uuid.UUID
	payment_integration_id uuid.UUID
	payment_status         varchar(10)
);


CREATE TABLE IF NOT EXISTS "payment_request" (
	id       		   uuid.UUID
    external_reference varchar(255),
	notification_url   varchar(255) not null,
	total_amount       int,
	expiration_date    timestamp
);