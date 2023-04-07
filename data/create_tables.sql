CREATE TABLE users(
   id SERIAL PRIMARY KEY,
   username VARCHAR(50) UNIQUE,
   firstname TEXT,
   lastname TEXT,
   email TEXT UNIQUE
);

CREATE TABLE species(
   id SERIAL PRIMARY KEY,
   name TEXT
);

CREATE TABLE advice(
   id SERIAL PRIMARY KEY,
   text TEXT,
   species_id INT NOT NULL REFERENCES species(id)
);

CREATE TABLE location(
   id SERIAL PRIMARY KEY,
   address VARCHAR(255),
   zip VARCHAR(5),
   city TEXT,
   lat DECIMAL(15,2),
   lon DECIMAL(15,2)
);

CREATE TABLE plants(
   id SERIAL PRIMARY KEY,
   location_id INT NOT NULL REFERENCES location(id),
   species_id INT NOT NULL REFERENCES species(id),
   user_id INT NOT NULL REFERENCES users(id)
);

CREATE TABLE photo(
   id SERIAL PRIMARY KEY,
   url TEXT,
   plant_id INT NOT NULL REFERENCES plants(id)
);

CREATE TABLE keeps(
   id SERIAL PRIMARY KEY,
   start_date DATE,
   end_date DATE,
   plant_id INT NOT NULL REFERENCES plants(id),
   user_id INT NOT NULL REFERENCES users(id),
   location_id INT NOT NULL REFERENCES location(id)
);

CREATE TABLE message(
   id SERIAL,
   sender_id INT NOT NULL REFERENCES users(id),
   receiver_id INT NOT NULL REFERENCES users(id),
   text VARCHAR(50),
   PRIMARY KEY(id, sender_id, receiver_id)
);

CREATE TABLE need_keep(
   user_id INT NOT NULL REFERENCES users(id),
   start_date DATE,
   end_date DATE,
   plant_id INT NOT NULL REFERENCES plants(id),
   PRIMARY KEY(user_id)
);
