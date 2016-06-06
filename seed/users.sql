DROP DATABASE IF EXISTS test;
CREATE DATABASE test;

\connect test

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL
);

INSERT INTO users (email, password) VALUES ('testuser@dto.gov.au', '$2a$10$vfN6rtYz9XWDoGVEFRx9fOVf39GIHjALBGzJhN5cidJhFAebcq85m'); -- pw = "password1"