CREATE TYPE roles_enum AS ENUM (
    'ADMIN',
    'DEFAULT',
    'COACH'
);



CREATE TYPE account_logs_tags_enum AS ENUM (
    'ANY',
    'BLOCK',
    'UNBLOCK',
    'CREATE',
    'DELETE',
    'UPDATE',
    'RECOVER',
    'SOFT_DELETE'
    'RESTORE'
);



CREATE TABLE account_tb(
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    birthdate TIMESTAMP NOT NULL,
    role ROLES_ENUM DEFAULT 'DEFAULT',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    soft_deleted BOOLEAN DEFAULT FALSE,
    is_blocked BOOLEAN DEFAULT FALSE,
    two_factor_code VARCHAR(6) DEFAULT '000000'
);



CREATE TABLE account_tb_logs(
    id SERIAL PRIMARY KEY,
    account_id INT REFERENCES account_tb(id),
    logged_at TIMESTAMP DEFAULT NOW(),
    tag ACCOUNT_LOGS_TAGS_ENUM DEFAULT 'ANY',
    msg VARCHAR(255) NOT NULL
);