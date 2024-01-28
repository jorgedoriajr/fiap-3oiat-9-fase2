CREATE TABLE IF NOT EXISTS "payment" (
	id       UUID not null primary key,
	order_id UUID not null references "order"(id),
	Data     bytea
);

CREATE TABLE IF NOT EXISTS "payment_status" (
	id                     UUID not null primary key,
	payment_id			   UUID references "payment"(id),
	payment_integration_id UUID,
	payment_status         varchar(10)
);


CREATE TABLE IF NOT EXISTS "payment_request" (
	id       		   UUID not null primary key,
    external_reference varchar(255),
	notification_url   varchar(255) not null,
	total_amount       int,
	expiration_date    timestamp
);