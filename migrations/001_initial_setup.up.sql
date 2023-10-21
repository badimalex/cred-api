-- migrations/001_initial_setup.up.sql
CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    surname VARCHAR(255),
    patronymic VARCHAR(255),
    age INTEGER,
    gender VARCHAR(10),
    nationality VARCHAR(255)
);
