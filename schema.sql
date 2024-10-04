CREATE TABLE users(
   id BIGSERIAL PRIMARY KEY,
   username text NOT NULL UNIQUE, 
   email text NOT NULL UNIQUE,
   password text NOT NULL,
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP
   
);