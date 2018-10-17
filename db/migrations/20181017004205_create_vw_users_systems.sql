-- migrate:up

CREATE VIEW vw_users_systems AS SELECT 
  S.id AS id, U.id AS user_id, U.user AS user, S.nombre AS nombre, 0 AS existe  
  FROM users_systems US 
  INNER JOIN users U ON U.id = US.user_id 
  INNER JOIN systems S ON S.id = US.system_id  
  LIMIT 2000;

-- migrate:down

DROP VIEW IF EXISTS vw_users_systems;
