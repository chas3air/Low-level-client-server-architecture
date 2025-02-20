CREATE DATABASE psql;

\c psql;

CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    role VARCHAR(20) NOT NULL,
    nick VARCHAR(50) NOT NULL
);

INSERT INTO Users (email, password, role, nick) VALUES  
('test@test.com', '123', 'user', 'nicK'),
('admin@admin.com', 'qwerty', 'admin', 'qaz');
