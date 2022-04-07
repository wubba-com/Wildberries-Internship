CREATE TABLE events (
    event_uid SERIAL PRIMARY KEY UNIQUE NOT NULL,
    name varchar (128) NOT NULL,
    date date NOT NULL
    );