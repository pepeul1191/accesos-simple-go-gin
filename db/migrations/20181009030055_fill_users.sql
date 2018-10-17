-- migrate:up

INSERT INTO users (user, pass, email, user_state_id) VALUES
  ('pepe', 'vPWKzwMiwQniX0HmpIrBW416CZzRDYksG991XW+a9iU=', 'jvaldivia@softweb.pe', 1),
  ('yacky', '+8IhO1fN6o4nlSfnNQOQzvMPg1QZtPwZsRsQXu1uSaE=', 'yramiez@softweb.pe', 2);

-- migrate:down

DELETE FROM users;
