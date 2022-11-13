```
docker pull postgres:14

docker run -e POSTGRES_PASSWORD=lol --name=pg --rm -d -p 5432:5432 postgres:14

docker exec -u postgres -it pg psql
```

```
# psql
\? show help
\d describe
\c change user
```

- Creating & Populating a table

```sql

CREATE DATABASE recipeguru;

CREATE TABLE ingredients (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  title VARCHAR ( 255 ) UNIQUE NOT NULL
);


INSERT INTO ingredients (title) VALUES ('bell pepper');

DROP TABLE ingredients;
```

- Altering a Table

```sql
CREATE TABLE ingredients (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  title VARCHAR ( 255 ) UNIQUE NOT NULL
);

ALTER TABLE ingredients ADD COLUMN image VARCHAR ( 255 );

ALTER TABLE ingredients DROP COLUMN image;

ALTER TABLE ingredients
ADD COLUMN image VARCHAR ( 255 ),
ADD COLUMN type VARCHAR ( 50 ) NOT NULL DEFAULT 'vegetable';
```

- Inserting Data and Mapping Conflict

```sql
INSERT INTO "ingredients" (
 "title", "image", "type" -- Notice the " here
) VALUES (
  'broccoli', 'broccoli.jpg', 'vegetable' -- and the ' here
);

INSERT INTO ingredients (
  title, image, type
) VALUES
  ( 'avocado', 'avocado.jpg', 'fruit' ),
  ( 'banana', 'banana.jpg', 'fruit' ),
  ( 'beef', 'beef.jpg', 'meat' ),
  ( 'black_pepper', 'black_pepper.jpg', 'other' ),
  ( 'blueberry', 'blueberry.jpg', 'fruit' ),
  ( 'broccoli', 'broccoli.jpg', 'vegetable' ),
  ( 'carrot', 'carrot.jpg', 'vegetable' ),
  ( 'cauliflower', 'cauliflower.jpg', 'vegetable' ),
  ( 'cherry', 'cherry.jpg', 'fruit' ),
  ( 'chicken', 'chicken.jpg', 'meat' ),
  ( 'corn', 'corn.jpg', 'vegetable' ),
  ( 'cucumber', 'cucumber.jpg', 'vegetable' ),
  ( 'eggplant', 'eggplant.jpg', 'vegetable' ),
  ( 'fish', 'fish.jpg', 'meat' ),
  ( 'flour', 'flour.jpg', 'other' ),
  ( 'ginger', 'ginger.jpg', 'other' ),
  ( 'green_bean', 'green_bean.jpg', 'vegetable' ),
  ( 'onion', 'onion.jpg', 'vegetable' ),
  ( 'orange', 'orange.jpg', 'fruit' ),
  ( 'pineapple', 'pineapple.jpg', 'fruit' ),
  ( 'potato', 'potato.jpg', 'vegetable' ),
  ( 'pumpkin', 'pumpkin.jpg', 'vegetable' ),
  ( 'raspberry', 'raspberry.jpg', 'fruit' ),
  ( 'red_pepper', 'red_pepper.jpg', 'vegetable' ),
  ( 'salt', 'salt.jpg', 'other' ),
  ( 'spinach', 'spinach.jpg', 'vegetable' ),
  ( 'strawberry', 'strawberry.jpg', 'fruit' ),
  ( 'sugar', 'sugar.jpg', 'other' ),
  ( 'tomato', 'tomato.jpg', 'vegetable' ),
  ( 'watermelon', 'watermelon.jpg', 'fruit' )
ON CONFLICT DO NOTHING;


INSERT INTO ingredients (
  title, image, type
) VALUES
  ( 'watermelon', 'banana.jpg', 'this won''t be updated' )
ON CONFLICT (title) DO UPDATE SET image = excluded.image;
```

- Updating & Deleting Data

```sql
UPDATE ingredients SET image = 'strawberry.jpg' WHERE title = 'watermelon' RETURN *;
UPDATE ingredients SET image = 'strawberry.jpg' WHERE title = 'watermelon' RETURNING *;


INSERT INTO ingredients
  (title, image, type)
VALUES
  ('not real 1', 'delete.jpg', 'nothing'),
  recipeguru=# SELECT recipe.title, recipe.body, photo.url
  FROM recipes_photos photo
  INNER JOIN
  recipes recipe
  ON
  photo.recipe_id = recipe.recipe_id;

  ('not real 2', 'delete.jpg', 'nothing');

DELETE FROM ingredients
WHERE image='delete.jpg'
RETURNING *;
```

- Selecing, Pagination & Using Where Clauses

```sql
recipeguru=# SELECT id, title, type FROM ingredients WHERE id > 24 LIMIT 10;
recipeguru=# SELECT id, title, type FROM ingredients LIMIT 10 OFFSET 30;
recipeguru=# SELECT id, title, type FROM ingredients WHERE type='fruit';
recipeguru=# SELECT id, title, type FROM ingredients WHERE type<>'fruit';  -- not equal
recipeguru=# SELECT id, title, type FROM ingredients WHERE type<>'fruit' AND id >=10 AND id <= 10 LIMIT 5;
recipeguru=# SELECT id, title, type FROM ingredients ORDER BY id DESC LIMIT 5;

```

 - LIKE, ILIKE & SQL [Functions](https://www.postgresql.org/docs/9.2/functions.html)

 ```sql
recipeguru=# SELECT * FROM ingredients WHERE title LIKE '%pota%';
recipeguru=# SELECT * FROM ingredients WHERE title ILIKE '%Pota%';
recipeguru=# SELECT * FROM ingredients WHERE CONCAT(title, type) ILIKE '%rryfru%';
recipeguru=# SELECT * FROM ingredients WHERE LOWER(CONCAT(title, type)) ILIKE '%rryfru%';

recipeguru=# SELECT * FROM ingredients WHERE title ILIKE 'ch_rry'; -- `_` means any one char
```

 - node-postgres & SQL Injection

 ```js
const pg = require("pg");
const pool = new pg.Pool({
  user: "postgres",
  host: "localhost",
  database: "recipeguru",
  password: "lol",
  port: 5432,
});


async function main() {
const { rows } = await pool.query(`SELECT * FROM ingredients`);
console.log(rows)
}

main()


// SQL injection
// id = 'SELECT * FROM ingredients WHERE id=1; DROP TABLE users; --'
const {id} = req.query;

// BAD
const {rows} = await pool.query(`SELECT * FROM ingredients WHERE id=${id}`);

// GOOD
const {rows} = await pool.query(`SELECT * FROM ingredients WHERE id=$1`, [id]);
// WHERE text ILIKE '%star wars%' => 'WHERE text ILIKE $1', ['%star wars%'])
 ```
 
 - Relationships & Joins

 ```sql

recipeguru=# CREATE TABLE recipes (
recipe_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
title VARCHAR (255) UNIQUE NOT NULL,
body TEXT
);

recipeguru=# INSERT INTO recipes
  (title, body)
VALUES
  ('cookies', 'very yummy'),
  ('empanada','ugh so good'),
  ('jollof rice', 'spectacular'),
  ('shakshuka','absolutely wonderful'),
  ('khachapuri', 'breakfast perfection'),
  ('xiao long bao', 'god I want some dumplings right now');
INSERT 0 6

recipeguru=# CREATE TABLE recipes_photos (
photo_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
recipe_id INTEGER,
url VARCHAR(255) NOT NULL
);

recipeguru=# INSERT INTO recipes_photos
  (recipe_id, url)
VALUES
  (1, 'cookies1.jpg'),
  (1, 'cookies2.jpg'),
  (1, 'cookies3.jpg'),
  (1, 'cookies4.jpg'),
  (1, 'cookies5.jpg'),
  (2, 'empanada1.jpg'),
  (2, 'empanada2.jpg'),
  (3, 'jollof1.jpg'),
  (4, 'shakshuka1.jpg'),
  (4, 'shakshuka2.jpg'),
  (4, 'shakshuka3.jpg'),
  (5, 'khachapuri1.jpg'),
  (5, 'khachapuri2.jpg');
-- no pictures of xiao long bao

```

- Left, Right & Outer Joins

```sql

recipeguru=# SELECT photo.recipe_id, photo.photo_id, recipe.title, recipe.body, photo.url
FROM recipes_photos photo
INNER JOIN
-- LEFT JOIN
-- RIGHT JOIN
-- RIGHT OUTER JOIN
-- FULL OUTER JOIN
recipes recipe
ON
photo.recipe_id = recipe.recipe_id
ORDER BY photo.photo_id;
```

- Natural & Cross Joins

```
recipeguru=# SELECT * FROM recipes NATURAL JOIN recipes_photos;

recipeguru=# SELECT r.title, r.body, rp.url 
FROM recipes_photos rp
CROSS JOIN recipes r;
```

- Foreign Keys & Managing References

```sql
recipeguru=# DELETE
FROM
recipes r
WHERE
r.recipe_id = 5;

recipeguru=# SELECT * FROM recipes_photos rp WHERE rp.recipe_id = 5;

# related photo left
```

```
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS recipes_photos;
CREATE TABLE recipes (
  recipe_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  title VARCHAR ( 255 ) UNIQUE NOT NULL,
  body TEXT
);
INSERT INTO recipes
  (title, body)
VALUES
  ('cookies', 'very yummy'),
  ('empanada','ugh so good'),
  ('jollof rice', 'spectacular'),
  ('shakshuka','absolutely wonderful'),
  ('khachapuri', 'breakfast perfection'),
  ('xiao long bao', 'god I want some dumplings right now');
```

```
recipeguru=# CREATE TABLE recipes_photos (
photo_id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
url VARCHAR(255) NOT NULL,
recipe_id INT REFERENCES recipes(recipe_id) ON DELETE CASCADE
);
```

- Many-to-Many Relationships

```
recipeguru=# CREATE TABLE recipe_ingredients (
recipe_id INT REFERENCES recipes(recipe_id) ON DELETE NO ACTION,
ingredient_id INT REFERENCES ingredients(id) ON DELETE NO ACTION,
CONSTRAINT recipe_ingredients_pk PRIMARY KEY (recipe_id, ingredient_id)

INSERT INTO recipe_ingredients
  (recipe_id, ingredient_id)
VALUES
  (1, 10),
  (1, 11),
  (1, 13),
  (2, 5),
  (2, 13);

SELECT
	i.title AS ingredient_title,
	i.image AS ingredient_image,
	i.type AS ingredient_type,
	r.title AS recipe_title,
	r.body AS recipe_body,
	r.recipe_id AS rid,
	i.id AS iid
FROM 
	recipe_ingredients ri 
INNER JOIN
	ingredients i
ON 
	i.id = ri.ingredient_id
INNER JOIN
	recipes r
ON
	r.recipe_id = ri.recipe_id
```

- Using the CHECK Constraint

```sql
recipeguru=# ALTER TABLE ingredients
ADD CONSTRAINT type_enums
CHECK
 (type IN('meat', 'fruit', 'vegetable', 'other'));

recipeguru=# INSERT INTO ingredients
(title, image, type)
VALUES ('hello', 'hello.jpg', 'hello');

-- ERROR:  new row for relation "ingredients" violates check constraint "type_enums"
-- DETAIL:  Failing row contains (70, hello, hello.jpg, hello).
```

- Using the DISTINCT Statement

```sql
recipeguru=# SELECT DISTINCT type FROM ingredients;

recipeguru=# SELECT DISTINCT ON (recipe_id) * FROM recipe_ingredients;

recipeguru=# SELECT DISTINCT ON (r.recipe_id)
	r.title,
	COALESCE(rp.url, 'default.jpg') AS url
FROM
	recipes r
LEFT JOIN
	recipes_photos rp
ON 
	r.recipe_id = rp.recipe_id;


SELECT DISTINCT ON (r.recipe_id)
	r.recipe_id, r.title, COALESCE(rp.url, 'default,jpg') AS url
FROM
	recipes r
LEFT JOIN
	recipes_photos rp
ON
	r.recipe_id = rp.recipe_id;
WHERE ri.recipe_id = 1
```
