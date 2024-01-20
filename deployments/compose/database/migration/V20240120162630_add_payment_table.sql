CREATE TABLE IF NOT EXISTS "payment" (
    external_reference varchar(255),
	notification_url   varchar(255) not null,
	total_amount    int,
	expiration_date time
);
