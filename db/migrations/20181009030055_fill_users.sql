-- migrate:up

INSERT INTO users (user, pass, email, user_state_id) VALUES
  ('pepe', 'kiki123', 'jvaldivia@softweb.pe', 1);

-- migrate:down

TRUNCATE users;
