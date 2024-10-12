-- name: SetSubscriber :one
INSERT INTO subscribers (id, uuid, email, name, attribs) 
VALUES ($1, gen_random_uuid(), $2, $3, $4)
ON CONFLICT (id) DO UPDATE 
  SET email = excluded.email, name = excluded.name, attribs = excluded.attribs, updated_at = CURRENT_TIMESTAMP
RETURNING *;

-- name: FindSubscriberByEmail :one
SELECT * FROM subscribers WHERE LOWER(email) = LOWER($1);

-- name: SubscribeList :one
INSERT INTO subscriber_lists (subscriber_id, list_id, status) 
VALUES ($1, $2, $3)
ON CONFLICT (subscriber_id, list_id) DO UPDATE 
  SET status = excluded.status
RETURNING *;