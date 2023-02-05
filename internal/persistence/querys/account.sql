-- SELECT * FROM pg_enum;
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


-- Sign Up
INSERT INTO account_tb(fullname, username, email, password, birthdate, role)
VALUES (
        'USER NAME',
        'username',
        'email@gmail.com',
        'senha_admin123',
        '2004-06-03',
        'DEFAULT'
);



-- Sign In
SELECT 1 FROM account_tb WHERE email = 'email@gmail.com' AND password = 'senha_admin123';



-- Recover password
UPDATE account_tb
SET two_factor_code = '123456'
WHERE id = 1;



-- Recover password validation
SELECT 1
FROM account_tb
WHERE
    two_factor_code = '123456' AND
    email = 'email@gmail.com';



-- Update password
UPDATE account_tb
SET password = 'new_password'
WHERE email = 'email@gmail.com';



-- Account details
SELECT
    id,
    fullname,
    username,
    email,
    password,
    birthdate,
    role,
    created_at,
    updated_at,
    soft_deleted,
    is_blocked,
    two_factor_code
FROM
    account_tb;




-- Update Account (user request)
UPDATE account_tb
SET
    fullname = 'Bla bla bla',
    username = 'Bla bla bla',
    email = 'Bla bla bla @gmail.com',
    password = 'Bla bla bla',
    birthdate = NOW(),
    role = 'COACH'
WHERE
    two_factor_code = '123456' AND
    email = 'email@gmail.com'
;



-- SoftDelete Account
UPDATE account_tb
SET soft_deleted = true
WHERE
    two_factor_code = '123456' AND
    email = 'email@gmail.com'
;



-- Restore Account
UPDATE account_tb
SET soft_deleted = false
WHERE
    two_factor_code = '123456' AND
    email = 'email@gmail.com'
;



-- Block Account
UPDATE account_tb
SET is_blocked = true
WHERE
    two_factor_code = '123456' AND
    email = 'email@gmail.com'
;



-- Unlock Account
UPDATE account_tb
SET is_blocked = false
WHERE
    two_factor_code = '123456' AND
    email = 'email@gmail.com'
;