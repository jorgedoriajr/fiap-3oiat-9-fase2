CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "customer" (
    cpf                 varchar(14) primary key,
    phone               varchar(15),
    name                varchar(255) not null ,
    email               varchar(255),
    opt_in_promotion    boolean default true,
    created_at          timestamp not null,
    updated_at          timestamp,
    active              boolean default true
);

CREATE TABLE IF NOT EXISTS "product_category" (
    name            varchar(255) primary key,
    accept_custom   boolean not null default false
);

CREATE TABLE IF NOT EXISTS "ingredient_type" (
    name    varchar(50) primary key
);

CREATE TABLE IF NOT EXISTS "ingredient_type_product_category" (
    id                  UUID not null primary key default public.uuid_generate_v4(),
    ingredient_type     varchar(50) references "ingredient_type"(name),
    optional            boolean,
    max_qtd             int,
    product_category    varchar(50) references "product_category"(name)
);

CREATE TABLE IF NOT EXISTS "ingredient" (
    id          UUID not null primary key default public.uuid_generate_v4(),
    number      serial,
    name        varchar(255) unique not null,
    amount      bigint not null,
    type        varchar(50) references "ingredient_type"(name),
    active      boolean default true
);


CREATE TABLE IF NOT EXISTS "product" (
    id                      UUID not null primary key default public.uuid_generate_v4(),
    number                  serial,
    name                    varchar(255) not null,
    amount                  bigint not null,
    description             text,
    category                varchar(50) references "product_category"(name) not null ,
    menu                    boolean not null,
    img_path                varchar(255) not null,
    created_at              timestamp not null,
    updated_at              timestamp,
    active                  boolean default true
);

CREATE TABLE IF NOT EXISTS "product_ingredient" (
    id                      UUID not null primary key default public.uuid_generate_v4(),
    product_id              UUID references "product"(id) not null,
    ingredient_id           UUID references "ingredient"(id) not null,
    quantity                int not null,
    amount                  bigint not null
);

CREATE TABLE IF NOT EXISTS "order" (
    id              UUID not null primary key default public.uuid_generate_v4(),
    customer_id     varchar(14) references "customer"(cpf),
    payment_id      UUID,
    takeAway        boolean not null default false,
    amount          bigint not null,
    status          varchar(50) not null ,
    created_at      timestamp not null,
    updated_at      timestamp
);

CREATE TABLE IF NOT EXISTS "order_product" (
    id                      UUID not null primary key default public.uuid_generate_v4(),
    product_id              UUID references "product"(id) not null,
    order_id                UUID references "order"(id) not null,
    quantity                int not null,
    amount                  bigint not null
);

CREATE TABLE IF NOT EXISTS "order_history" (
    id          UUID not null primary key default public.uuid_generate_v4(),
    order_id    UUID references "order"(id) not null,
    status      varchar(50) not null,
    change_by   varchar(50) not null,
    created_at  timestamp not null
);


ALTER TABLE "order" ADD CONSTRAINT fk_order_customer
    FOREIGN KEY (customer_id) REFERENCES "customer"(cpf);

ALTER TABLE "order_history" ADD CONSTRAINT fk_order_history_order
    FOREIGN KEY (order_id) REFERENCES "order"(id);

CREATE UNIQUE INDEX IF NOT EXISTS product_number_unq_idx ON product(number);
CREATE UNIQUE INDEX IF NOT EXISTS ingredient_number_unq_idx ON ingredient(number);
