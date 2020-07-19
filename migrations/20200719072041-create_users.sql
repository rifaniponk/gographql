
-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    username VARCHAR (127) NOT NULL UNIQUE,
    email VARCHAR (127) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS links(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    title VARCHAR (255),
    address VARCHAR (255),
    userid UUID
);

-- +migrate Down
DROP TABLE users;
DROP TABLE links;
