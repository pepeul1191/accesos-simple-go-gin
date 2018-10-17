-- migrate:up

CREATE TABLE systems (
	id	INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	nombre	VARCHAR(30) NOT NULL,
	version	VARCHAR(6),
	repositorio	VARCHAR(200)
);

-- migrate:down

DROP TABLE IF EXISTS systems;
