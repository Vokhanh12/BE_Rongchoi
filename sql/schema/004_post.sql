-- +goose Up
CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(255) NOT NULL,
    number_phone VARCHAR(20) NOT NULL,
    address VARCHAR(255) NOT NULL,
    nick_name VARCHAR(15) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
    
);
-- +goose Down
DROP TABLE posts;
