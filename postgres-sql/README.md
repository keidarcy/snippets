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

- JSONB

```sql
recipeguru=# ALTER TABLE recipes
ADD COLUMN meta JSONB;

recipeguru=# UPDATE recipes
SET
meta = '{ "tag": ["chocolate", "dessert", "cake"] }'
WHERE
recipe_id=16;
UPDATE 1

UPDATE
  recipes
SET
  meta='{ "tags": ["dessert", "cake"] }'
WHERE
  recipe_id=20;

UPDATE
  recipes
SET
  meta='{ "tags": ["dessert", "fruit"] }'
WHERE
  recipe_id=45;

UPDATE
  recipes
SET
  meta='{ "tags": ["dessert", "fruit"] }'
WHERE
  recipe_id=47;

```

```sql
recipeguru=# SELECT meta, recipe_id  FROM recipes WHERE meta IS NOT NULL;

recipeguru=# SELECT meta -> 'tags' -> 0 AS tag, recipe_id  FROM recipes WHERE meta IS NOT NULL;

recipeguru=# SELECT meta -> 'tags' ->> 0 AS tag, recipe_id  FROM recipes WHERE meta IS NOT NULL; -- get string

recipeguru=# SELECT recipe_id, title, meta -> 'tags' AS has_cake_tag FROM recipes
WHERE meta -> 'tags' ? 'cake'; -- does tags contains 'cake';

recipeguru=# SELECT recipe_id, title, meta -> 'tags' AS has_cake_tag FROM recipes
WHERE meta -> 'tags' @> '"cake"';
```

- Aggregation

```sql
recipeguru=# SELECT COUNT(DISTINCT type), type  FROM ingredients;

recipeguru=# SELECT COUNT(DISTINCT type) FROM ingredients;

recipeguru=# SELECT
    type, COUNT(type)
FROM
    ingredients
GROUP BY
    type;

```

- Filtering Aggregation

```sql
recipeguru=# SELECT
type, COUNT(type)
FROM
ingredients
WHERE
image IS NOT NULL
GROUP BY
type
HAVING COUNT(type) > 10;
```

- Storing Queries with Functions

```sql
SELECT
  r.title
FROM recipe_ingredients ri

INNER JOIN
  recipes r
ON
  r.recipe_id = ri.recipe_id

GROUP BY
  r.title
HAVING
  COUNT(r.title) BETWEEN 4 AND 6


CREATE OR REPLACE FUNCTION
  get_recipes_with_ingredients(low INT, high INT)
RETURNS
  SETOF VARCHAR
LANGUAGE
  plpgsql
AS
$$
BEGIN
  RETURN QUERY SELECT
    r.title
  FROM
    recipe_ingredients ri

  INNER JOIN
    recipes r
  ON
    r.recipe_id = ri.recipe_id
  GROUP BY
    r.title
  HAVING
    COUNT(r.title) between low and high;
END;
$$;

SELECT * FROM get_recipes_with_ingredients(1, 1);

\df
```

- Functions vs Procedures

```sql
SELECT * FROM ingredients WHERE images IS NULL;

INSERT INTO ingredients (title, type) VALUES ('version', 'meat');



CREATE PROCEDURE
  set_null_ingredient_images_to_default()
LANGUAGE
  SQL
AS
$$
  UPDATE
    ingredients
  SET
    image = 'default.jpg'
  WHERE
    image IS NULL;
$$;

CALL set_null_ingredient_images_to_default();
```

- Calling Functions with Triggers

```sql
recipeguru=# CREATE TABLE updated_recipes (
  id INT GENERATED ALWAYS AS IDENTITY,
  recipe_id INT,
  old_title VARCHAR(255),
  new_title VARCHAR(255),
  time_of_update TIMESTAMP
);

CREATE OR REPLACE FUNCTION log_updated_recipe_name()
  RETURNS
    TRIGGER
  LANGUAGE
    plpgsql
AS
$$
BEGIN
  IF OLD.title <> NEW.title THEN
    INSERT INTO
      updated_recipes (recipe_id, old_title, new_title, time_of_update)
    VALUES
      (NEW.recipe_id, OLD.title, NEW.title, NOW());
  END IF;
  RETURN NEW;
END;
$$;


CREATE OR REPLACE TRIGGER updated_recipe_trigger
AFTER UPDATE ON recipes
  FOR EACH ROW EXECUTE PROCEDURE log_updated_recipe_name();
```

- The Movie Database

```sql
-- How much revenue did the movies Keanu Reeves act in make?
SELECT
  SUM(revenue) AS total
FROM
  movies m

JOIN
  casts c
ON
  m.id = c.movie_id

JOIN
  people p
ON
  c.person_id = p.id
WHERE
  p.name = 'Keanu Reeves';
```

```sql
-- Which 5 people were in the movies that had the most revenue?
SELECT
  p.name, SUM(COALESCE(revenue, 0)) AS money
FROM
  movies m

JOIN
  casts c
ON
  c.movie_id = m.id

JOIN
  people p
ON
  c.person_id = p.id
GROUP BY
p.name
ORDER BY
  money DESC

LIMIT
5;
```

```sql
-- Which 10 movies have the most keywords?
SELECT
  m.name AS movie_name, COUNT(c.id) AS keywords_number
FROM
  movies m

JOIN
  movie_keywords k
ON
  m.id = k.movie_id

JOIN
  categories c
ON
  c.id = k.category_id
GROUP BY
  movie_name
ORDER BY
  keywords_number DESC

LIMIT
  10;
```

```sql
-- Which category is associated with the most movies
SELECT
  c.name AS keyword, COUNT(c.id) AS keyword_count
FROM
  categories c

JOIN
  movie_keywords k
ON
  k.category_id = c.id

GROUP BY
  keyword, k.category_id
ORDER BY
  keyword_count DESC
LIMIT
  10;

```

- Analyzing Queries with EXPLAIN

```sql
SELECT * FROM movies WHERE name='Tron Legacy';
SELECT * FROM movies WHERE id=21103;

EXPLAIN SELECT * FROM movies WHERE name='Tron Legacy';
EXPLAIN SELECT * FROM movies WHERE id=21103;

EXPLAIN ANALYZE SELECT * FROM movies WHERE name='Tron Legacy';
EXPLAIN ANALYZE SELECT * FROM movies WHERE id=21103;
```


- Creating an Index

```sql
omdb=# EXPLAIN ANALYZE SELECT * FROM movies WHERE name='Tron Legacy';
                                              QUERY PLAN
-------------------------------------------------------------------------------------------------------
 Seq Scan on movies  (cost=0.00..5060.27 rows=2 width=75) (actual time=42.193..430.487 rows=1 loops=1)
   Filter: (name = 'Tron Legacy'::text)
   Rows Removed by Filter: 178395
 Planning Time: 4.741 ms
 Execution Time: 430.847 ms
(5 rows)

omdb=# CREATE INDEX idx_name ON movies(name);
CREATE INDEX
omdb=# EXPLAIN ANALYZE SELECT * FROM movies WHERE name='Tron Legacy';
                                                   QUERY PLAN
-----------------------------------------------------------------------------------------------------------------
 Bitmap Heap Scan on movies  (cost=4.44..12.30 rows=2 width=75) (actual time=2.393..2.532 rows=1 loops=1)
   Recheck Cond: (name = 'Tron Legacy'::text)
   Heap Blocks: exact=1
   ->  Bitmap Index Scan on idx_name  (cost=0.00..4.43 rows=2 width=0) (actual time=1.065..1.067 rows=1 loops=1)
         Index Cond: (name = 'Tron Legacy'::text)
 Planning Time: 5.747 ms
 Execution Time: 4.750 ms
(7 rows)

omdb=# DROP INDEX idx_name;
DROP INDEX
```

- Using a GIN Index(generalized inverted index)

```sql
CREATE TABLE movie_reviews (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  movie_id INTEGER,
  scores JSONB NOT NULL
);

INSERT INTO
  movie_reviews
  (movie_id, scores)
VALUES
  (21103, '{ "rotten_tomatoes": 94, "washington_post": 50, "nytimes": 45 }'),
  (97, '{ "rotten_tomatoes": 87, "washington_post": 40 }'),
  (18235, '{ "rolling_stone": 85, "washington_post": 60, "nytimes": 35 }'),
  (10625, '{ "rotten_tomatoes": 100, "washington_post": 100, "nytimes": 100, "rolling_stone": 100 }'),
  (85014, '{ "nytimes": 67 }'),
  (2493, '{ "rotten_tomatoes": 67, "rolling_stone": 89, "nytimes": 85 }'),
  (11362, '{ "rotten_tomatoes": 76, "washington_post": 14, "nytimes": 98 }'),
  (674, '{ "rotten_tomatoes": 78, "washington_post": 40, "nytimes": 77, "rolling_stone": 54 }');

CREATE INDEX ON movie_reviews USING gin(scores);

EXPLAIN ANALYZE SELECT * FROM movie_reviews WHERE scores @> '{"nytimes": 98}';

omdb=# SELECT SHOW_TRGM('star wars');

omdb=# CREATE INDEX ON movies USING gin(name gin_trgm_ops);
```

- Creating a Partial Index

```sql
SELECT DISTINCT language, COUNT(*) FROM category_names GROUP BY language;

CREATE INDEX idx_en_category_names ON category_names(language) WHERE language='en';

EXPLAIN ANALYZE SELECT * FROM category_names WHERE language='ge' AND name ILIKE '%animation%' LIMIT 5;
```

- Indexing a Derivative Value;

```sql
EXPLAIN ANALYZE SELECT
  name, date, revenue, budget, COALESCE((revenue - budget), 0) AS profit
FROM
  movies
ORDER BY
  profit DESC
LIMIT 10;


CREATE INDEX idx_movies_profit ON movies (COALESCE((revenue - budget), 0));
```

- Creating Views & Inserting Data

```sql
CREATE VIEW english_category_names AS SELECT category_id, name, language FROM category_names WHERE language='en' LIMIT 5;

SELECT * FROM english_category_names LIMIT 15;
```
