CREATE DATABASE blog_service;

CREATE OR REPLACE FUNCTION uuid_generate_v4()
RETURNS uuid
AS $$
BEGIN
  RETURN ('' || md5(random()::text || clock_timestamp()::text))::uuid;
END;
$$
LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE blogs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    blog VARCHAR(255),
    like INT,
    comment VARCHAR(255),
    author VARCHAR(255),
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE likes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    blog_id UUID REFERENCES blogs(id),
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    blog_id UUID REFERENCES blogs(id),
    user_id UUID REFERENCES users(id),
    text VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);


DROP DATABASE blog_service;