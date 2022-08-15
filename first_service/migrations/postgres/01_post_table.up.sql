
CREATE TABLE IF NOT EXISTS posts (
    id INT NOT NULL,
    user_id INT NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at  timestamp DEFAULT (now()),
    updated_at timestamp,
    deleted_at timestamp DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS check_post (
    message TEXT NOT NULL,
    error TEXT NOT NULL,
    created_at timestamp DEFAULT (now())
);