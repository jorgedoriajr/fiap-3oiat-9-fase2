-- product_category
INSERT INTO product_category(name, accept_custom)
VALUES ('Bebida', false);

INSERT INTO product_category(name, accept_custom)
VALUES ('Acompanhamento', false);

INSERT INTO product_category(name, accept_custom)
VALUES ('Lanche', true);

-- ingredient_type
INSERT INTO ingredient_type(name)
VALUES('Pão');

INSERT INTO ingredient_type(name)
VALUES('Hamburger');

INSERT INTO ingredient_type(name)
VALUES('Queijo');

INSERT INTO ingredient_type(name)
VALUES('Molho');

INSERT INTO ingredient_type(name)
VALUES('Adicionais');

INSERT INTO ingredient_type(name)
VALUES('Acompanhamento');

INSERT INTO ingredient_type(name)
VALUES('Bebidas Alcoólicas');

INSERT INTO ingredient_type(name)
VALUES('Bebidas Não Alcoólicas');

-- ingredient_type_product_category
INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Pão', false, 1, 'Lanche');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Hamburger', false, 2, 'Lanche');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Queijo', true, 2, 'Lanche');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Queijo', true, 1, 'Acompanhamento');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Molho', true, null, 'Lanche');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Molho', true, null, 'Acompanhamento');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Adicionais', true, null, 'Lanche');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Adicionais', true, null, 'Acompanhamento');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Acompanhamento', true, null, 'Acompanhamento');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Bebidas Alcoólicas', false, null, 'Bebida');

INSERT INTO ingredient_type_product_category(ingredient_type, optional, max_qtd, product_category)
VALUES ('Bebidas Não Alcoólicas', false, null, 'Bebida');

-- Ingredient
INSERT INTO ingredient(id, name, amount, type)
VALUES ('ba97de21-2721-40cc-a326-f3bfebedf862', 'Com Gergelim', 500, 'Pão');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('869c0abf-2c16-4df3-a24b-a5c15844ef56', 'Sem Gergelim', 500, 'Pão');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('ef045802-2937-4b0d-a2c5-938b140ba4c3', 'Carne Angus', 2500, 'Hamburger');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('b397d83f-02cf-4e2a-ac35-36c87910b2f5', 'Vegano', 2500, 'Hamburger');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('ea7c6b6a-f8b5-4305-89d3-d6d22d85f9b0','Gorgonzola', 600, 'Queijo');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('b9cae298-e74a-4bc7-985d-d268f2bc76a6', 'Cheddar', 500, 'Queijo');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('6adddd70-ab19-4fc2-a50e-cefc5ab7605b', 'Maionese da Casa', 500, 'Molho');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('10bfa93e-e061-4dfc-8302-95b5a010c7b6', 'Maionese Picante', 500, 'Molho');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('d4d14983-d431-48aa-9e2c-8a14e704101d', 'Bacon', 600, 'Adicionais');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('ad2f5251-5053-451a-808e-2985a5940216', 'Picles', 300, 'Adicionais');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('6c51f438-a55e-426e-98f2-c1bd03e70b3e', 'Coca-Cola Lata', 700, 'Bebidas Não Alcoólicas');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('501d1ad7-f002-446c-81ce-59b75f6599ec', 'Heineken Long Neck', 1000, 'Bebidas Alcoólicas');

INSERT INTO ingredient(id, name, amount, type)
VALUES ('85f0288e-f692-458e-ab0f-93c31adb5579', 'Batata', 2000, 'Acompanhamento');

-- product
INSERT INTO product(id, name, amount, description, category, menu, created_at, updated_at)
VALUES (
    '787b1e5d-c5d8-4ed4-9554-ec0e6555c8f1',
    'Burger Cheddar Bacon', 4100, 'Burger blend de 220g de Angus, cheddar e bacon', 'Lanche', true, now(), now()
);

INSERT INTO product(id, name, amount, description, category, menu, created_at, updated_at)
VALUES (
   'ec2e5922-f616-48b4-9e46-97f086db3e7f',
   'Double Cheddar Burger', 6500, '2 Burgers blend de 220g de Angus e duplo cheddar', 'Lanche', true, now(), now()
);

INSERT INTO product(id, name, amount, description, category, menu, created_at, updated_at)
VALUES (
    '013c25c4-fc72-4919-af03-4379aff7a989',
    'Burger Vegano', 3000, 'Burguer blend de 220g vegano', 'Lanche', true, now(), now()
);

INSERT INTO product(id, name, amount, description, category, menu, created_at, updated_at)
VALUES (
   'ee491b9b-36d4-4b57-9aa6-21df5885c21e',
   'Heineken Long Neck', 1000, 'Heineken Long Neck 330ml', 'Bebida', true, now(), now()
);

INSERT INTO product(id, name, amount, description, category, menu, created_at, updated_at)
VALUES (
   '9724f90f-c557-4a31-a69c-57c7ddb91e69',
   'Coca-Cola Lata', 1000, 'Coca-Cola Lata 350ml', 'Bebida', true, now(), now()
);

INSERT INTO product(id, name, amount, description, category, menu, created_at, updated_at)
VALUES (
   '343ffe88-8d87-41f0-9433-a212d48a7c0b',
   'Batata com cheddar e bacon', 3100, 'Porção de batata com cheddar e bacon', 'Acompanhamento', true, now(), now()
);

-- product_ingredient

-- 'Batata com cheddar e bacon'
INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('343ffe88-8d87-41f0-9433-a212d48a7c0b', '85f0288e-f692-458e-ab0f-93c31adb5579', 1, 2000)

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('343ffe88-8d87-41f0-9433-a212d48a7c0b', 'b9cae298-e74a-4bc7-985d-d268f2bc76a6', 1, 500)

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('343ffe88-8d87-41f0-9433-a212d48a7c0b', 'd4d14983-d431-48aa-9e2c-8a14e704101d', 1, 600)

-- 'Coca-Cola Lata'
INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('9724f90f-c557-4a31-a69c-57c7ddb91e69', '6c51f438-a55e-426e-98f2-c1bd03e70b3e', 1, 700);

-- 'Heineken Long Neck'
INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('ee491b9b-36d4-4b57-9aa6-21df5885c21e', '501d1ad7-f002-446c-81ce-59b75f6599ec', 1, 1000);

-- 'Burger Cheddar Bacon'
INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('787b1e5d-c5d8-4ed4-9554-ec0e6555c8f1', 'ba97de21-2721-40cc-a326-f3bfebedf862', 1, 500);

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('787b1e5d-c5d8-4ed4-9554-ec0e6555c8f1', 'ef045802-2937-4b0d-a2c5-938b140ba4c3', 1, 2500);

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('787b1e5d-c5d8-4ed4-9554-ec0e6555c8f1', 'b9cae298-e74a-4bc7-985d-d268f2bc76a6', 1, 500);

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('787b1e5d-c5d8-4ed4-9554-ec0e6555c8f1', 'd4d14983-d431-48aa-9e2c-8a14e704101d', 1, 600);

-- 'Double Cheddar Burger'
INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('ec2e5922-f616-48b4-9e46-97f086db3e7f', 'ba97de21-2721-40cc-a326-f3bfebedf862', 1, 500);

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('ec2e5922-f616-48b4-9e46-97f086db3e7f', 'ef045802-2937-4b0d-a2c5-938b140ba4c3', 2, 5000);

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('ec2e5922-f616-48b4-9e46-97f086db3e7f', 'b9cae298-e74a-4bc7-985d-d268f2bc76a6', 2, 1000);

-- 'Burger Vegano'
INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('013c25c4-fc72-4919-af03-4379aff7a989', 'ba97de21-2721-40cc-a326-f3bfebedf862', 1, 500);

INSERT INTO product_ingredient(product_id, ingredient_id, quantity, amount)
VALUES ('013c25c4-fc72-4919-af03-4379aff7a989', 'b397d83f-02cf-4e2a-ac35-36c87910b2f5', 1, 2500);

