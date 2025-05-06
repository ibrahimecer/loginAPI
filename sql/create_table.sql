CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY NOT NULL,
    role_name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    user_surname VARCHAR(50) NOT NULL,
    user_nickname VARCHAR(75) UNIQUE NOT NULL,
    user_email VARCHAR(75) UNIQUE NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    role_id INTEGER NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY NOT NULL,
    exam1 INTEGER NOT NULL,
    exam2 INTEGER NOT NULL,
    exam3 INTEGER NOT NULL,
    average FLOAT,
    student_id INTEGER NOT NULL,
    FOREIGN KEY (student_id) REFERENCES users(id)
);
