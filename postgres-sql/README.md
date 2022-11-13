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

 ```
