-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follow.*, feeds.name, users.name AS user_name
FROM inserted_feed_follow
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name, users.name AS user_name FROM feed_follows
JOIN feeds ON feeds.id = feed_follows.feed_id
JOIN users on users.id = feed_follows.user_id
WHERE users.name = $1;
