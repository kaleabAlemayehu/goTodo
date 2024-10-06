CREATE TABLE users(
   id BIGSERIAL PRIMARY KEY,
   username text NOT NULL UNIQUE, 
   email text NOT NULL UNIQUE,
   password text NOT NULL,
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP
   
);

CREATE TABLE todos(
   id BIGSERIAL PRIMARY KEY,
   user_id INT REFERENCES users(id),
   title text NOT NULL,
   content text NOT NULL,
   starting_time timestamp,
   ending_time timestamp,
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);