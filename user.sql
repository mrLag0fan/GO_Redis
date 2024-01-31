\c "go-user-db"
DROP TABLE IF EXISTS "go-user";
CREATE TABLE "go-user" (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    email VARCHAR,
    age INTEGER
);
