CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users
(
  id int AUTO_INCREMENT PRIMARY KEY,
  name varchar(50) NOT NULL,
  nick varchar(50) NOT NULL,
  email varchar(50) NOT NULL,
  password varchar(50) NOT NULL,
  created_at timestamp default current_timestamp()
) ENGINE=InnoDB;