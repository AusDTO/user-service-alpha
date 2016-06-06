DROP DATABASE IF EXISTS test;
CREATE DATABASE test;

\connect test

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL
);

INSERT INTO users (email, password) VALUES ('testuser@dto.gov.au', 'abc');