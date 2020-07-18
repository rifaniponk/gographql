
-- +migrate Up
CREATE TABLE IF NOT EXISTS Users(
    ID UUID NOT NULL UNIQUE,
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Links(
    ID INT NOT NULL UNIQUE,
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID UUID
);

-- +migrate Down
