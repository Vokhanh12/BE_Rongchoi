-- +goose Up
CREATE TYPE user_role AS ENUM ('buyer', 'employee');

CREATE TABLE users(
    id UUID PRIMARY KEY,
    create_at DATE NOT NULL,
    update_at DATE NOT NULL,
    first_name VARCHAR(10) NOT NULL,
    last_name VARCHAR(10) NOT NULL,
    email VARCHAR(50) NOT NULL,
    nick_name VARCHAR(15),
    number_phone VARCHAR(20) NOT NULL,
    day_of_birth DATE,
    address VARCHAR(255),
    role user_role NOT NULL
);


-- +goose Down
DROP TABLE users;
DROP TYPE user_role;
