-- migrate:up

CREATE TABLE user_states (
	id	INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name	VARCHAR(10) NOT NULL
);

-- migrate:down

DROP TABLE IF EXISTS user_states;
