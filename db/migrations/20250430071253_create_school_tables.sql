-- +goose Up
CREATE DATABASE IF NOT EXISTS db_inventory;

USE db_inventory;

-- Enable the UUID function for MySQL
SET @@global.sql_mode = 'STRICT_TRANS_TABLES';

-- Table Login
CREATE TABLE Login (
      id INT AUTO_INCREMENT PRIMARY KEY,
        id2 CHAR(36), -- Store UUIDs as CHAR(36)
        username VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        account_status VARCHAR(255) NOT NULL,
        account_group VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        is_online BOOLEAN DEFAULT FALSE
);

-- Table Login
CREATE TABLE inventory (
      id INT AUTO_INCREMENT PRIMARY KEY,
      item_no VARCHAR(255) NOT NULL UNIQUE,
      item_name VARCHAR(255) NOT NULL,
      company VARCHAR(255) NOT NULL,
      date_entry VARCHAR(255) NOT NULL,
      available BOOLEAN DEFAULT FALSE,
      date_log TIMESTAMP DEFAULT CURRENT_TIMESTAMP
      
);

-- Table: teachers
-- CREATE TABLE teachers (
--                           id CHAR(36) PRIMARY KEY, -- Store UUIDs as CHAR(36)
--                           first_name VARCHAR(255) NOT NULL,
--                           last_name VARCHAR(255) NOT NULL,
--                           subject VARCHAR(255) NOT NULL,
--                           class VARCHAR(255) NOT NULL,
--                           email VARCHAR(255) NOT NULL UNIQUE
-- );

-- Table: students
-- CREATE TABLE students (
--                           id CHAR(36) PRIMARY KEY, -- Store UUIDs as CHAR(36)
--                           first_name VARCHAR(255) NOT NULL,
--                           last_name VARCHAR(255) NOT NULL,
--                           class VARCHAR(255) NOT NULL,
--                           email VARCHAR(255) NOT NULL UNIQUE
-- );

-- Table: executives
-- CREATE TABLE executives (
--                             id CHAR(36) PRIMARY KEY, -- Store UUIDs as CHAR(36)
--                             first_name VARCHAR(255) NOT NULL,
--                             last_name VARCHAR(255) NOT NULL,
--                             email VARCHAR(255) NOT NULL UNIQUE,
--                             username VARCHAR(255) NOT NULL UNIQUE,
--                             password CHAR(60) NOT NULL, -- To store bcrypt-hashed passwords
--                             last_password_change TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--                             user_creation_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--                             password_reset_token VARCHAR(255),
--                             reset_token_expiry TIMESTAMP,
--                             role VARCHAR(255) NOT NULL,
--                             user_inactive BOOLEAN DEFAULT FALSE
-- );
