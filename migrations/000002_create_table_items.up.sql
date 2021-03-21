CREATE TABLE items (
	id INT NOT NULL primary key,
	name TEXT NOT NULL,
	location TEXT,
	counts INT NOT NULL,
	manager_id INT NOT NULL
);
