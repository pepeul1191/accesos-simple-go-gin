-- migrate:up

INSERT INTO user_states (name) VALUES 
  ('active'),
  ('inactive'),
  ('suspended'),
  ('deleted'),
  ('vacationing');

-- migrate:down

TRUNCATE user_states;
