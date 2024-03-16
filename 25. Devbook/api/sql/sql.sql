/* MYSQL CLI
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    userpass VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;
*/

/*
Após realizar a instalção do postgres proceder com os scripts a seguir
para criação de novo user e garantir privilégios

CREATE USER golang WITH PASSWORD 'golang';
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO golang WITH GRANT OPTION;
SELECT * FROM information_schema.table_privileges WHERE grantee = 'golang';
REVOKE ALL PRIVILEGES ON usuarios FROM golang;
SELECT * FROM information_schema.table_privileges WHERE grantee = 'golang';
*/


/*
INSERT INTO usuarios (nome, email) VALUES ('Nome do Usuário', 'email@exemplo.com');
DELETE FROM usuarios;
ALTER SEQUENCE usuarios_id_seq RESTART WITH 1;
*/



/* POSTGRES CLI
CREATE DATABASE IF NOT EXISTS devbook;
\c devbook;
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    userpass VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
*/
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS followers CASCADE;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    userpass VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE followers (
    user_id INT NOT NULL,
    follower_id INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_follower
        FOREIGN KEY(follower_id) 
        REFERENCES users(id)
        ON DELETE CASCADE,
    PRIMARY KEY (user_id, follower_id)
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    likes INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    author_id INT NOT NULL,
    CONSTRAINT fk_post
        FOREIGN KEY(author_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
);