CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE
);

CREATE TABLE species(id SERIAL PRIMARY KEY, name TEXT);

CREATE TABLE advice(
    id SERIAL PRIMARY KEY,
    text TEXT,
    species_id INT NOT NULL,
    FOREIGN KEY (species_id) REFERENCES species(id)
);

CREATE TABLE location(
    id SERIAL PRIMARY KEY,
    address VARCHAR(255),
    zip VARCHAR(5),
    city TEXT,
    lat DECIMAL(15, 2),
    lon DECIMAL(15, 2)
);

CREATE TABLE plants(
    id SERIAL PRIMARY KEY,
    location_id INT NOT NULL,
    species_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES location(id),
    FOREIGN KEY (species_id) REFERENCES species(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE keeps(
    user_id INT,
    location_id INT,
    PRIMARY KEY (user_id, location_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (location_id) REFERENCES location(id)
);

CREATE TABLE message(
    sender_id INT,
    recipient_id INT,
    text VARCHAR(50),
    PRIMARY KEY (sender_id, recipient_id),
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (recipient_id) REFERENCES users(id)
);