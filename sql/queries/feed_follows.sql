-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follows.*, users.name as user_name, feeds.name as feed_name
FROM inserted_feed_follows
JOIN users on users.id = inserted_feed_follows.user_id 
JOIN feeds on feeds.id = inserted_feed_follows.feed_id;

-- name: GetFeedFollowsUser :many
SELECT *, users.name as user_name, feeds.name as feed_name
FROM feed_follows
JOIN users on users.id = feed_follows.user_id
JOIN feeds on feeds.id = feed_follows.feed_id
WHERE users.name = $1;