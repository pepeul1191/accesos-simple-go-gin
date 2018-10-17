-- migrate:up

CREATE VIEW vw_users_systems AS SELECT 
  US.id, S.id AS system_id, S.name AS system, U.id AS user_id, U.user AS user, U.pass AS pass 
  FROM users_systems US 
  INNER JOIN users U ON U.id = US.user_id 
  INNER JOIN systems S ON S.id = US.system_id  
  LIMIT 2000;

-- migrate:down

DROP VIEW IF EXISTS vw_users_systems;
