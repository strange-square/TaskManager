INSERT INTO projects(name, created_at, updated_at)
SELECT 
 md5(random()::text), 
 NOW() - '1 day'::INTERVAL * (RANDOM()::int * 100),
 NOW() - '1 day'::INTERVAL * (RANDOM()::int * 100)
FROM generate_series(1,50); 

INSERT INTO tasks(name, project_id, created_at, updated_at)
SELECT 
 md5(random()::text),
 id,
 NOW() - '1 day'::INTERVAL * (RANDOM()::int * 100), 
 NOW() - '1 day'::INTERVAL * (RANDOM()::int * 100)
FROM generate_series(1,50) id;