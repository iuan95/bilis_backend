CREATE TABLE IF NOT EXISTS posts(
   id SERIAL PRIMARY KEY,
   title VARCHAR (60) NOT NULL,
   description VARCHAR (250) NOT NULL,
   date TIMESTAMP NULL
);
