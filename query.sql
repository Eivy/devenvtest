CREATE TABLE users (
	id INT NOT NULL primary key,
	name TEXT NOT NULL
);

CREATE TABLE items (
	id INT NOT NULL primary key,
	name TEXT NOT NULL,
	location TEXT,
	counts INT NOT NULL,
	manager_id INT NOT NULL
);

-- name: GetUser :one
SELECT * FROM users where id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users order by name;

-- name: CreateUser :one
INSERT INTO users (
	id,
	name
)
SELECT
	CASE WHEN MAX(id) IS NULL THEN 0 ELSE MAX(id) + 1 END as id,
	$1 as name
FROM users
RETURNING *;

-- name: GetItem :one
SELECT * FROM items
LEFT JOIN users
ON items.manager_id = users.id
WHERE items.id = $1 LIMIT 1;

-- name: ListItems :many
SELECT * FROM items
LEFT JOIN users
ON items.manager_id = users.id
ORDER by items.name;

-- name: CreateItem :one
INSERT INTO items (
	id,
	name,
	location,
	counts,
	manager_id
)
SELECT
	CASE WHEN MAX(id) IS NULL THEN 0 ELSE MAX(id) + 1 END as id,
	$1 as name,
	$2 as location,
	$3 as counts,
	$4 as manager_id
FROM items
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;

-- name: DeleteItem :one
DELETE FROM items
WHERE id = $1
RETURNING *;
