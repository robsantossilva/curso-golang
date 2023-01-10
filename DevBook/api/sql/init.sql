CREATE TABLE users (
  id integer PRIMARY KEY,
  name varchar(50) NOT NULL,
  nick varchar(50) NOT NULL,
  email varchar(50) NOT NULL,
  password varchar(50) NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)