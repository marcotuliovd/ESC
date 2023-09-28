CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="America/Sao_Paulo";

-- Create  table
CREATE TABLE public.clients (
    id serial PRIMARY KEY,
    username varchar(255),
    name varchar(255),
    phone varchar(15),
    created_at timestamp
);

CREATE TABLE public.vehicle (
    id serial PRIMARY KEY,
    owner varchar(255),
    title varchar(255),
    type varchar(255),
    placa varchar(10),
    created_at timestamp
);

CREATE TABLE public.space (
    id serial PRIMARY KEY,
    vehicle_id integer,
    type varchar(255),
    created_at timestamp,
    FOREIGN KEY (vehicle_id) REFERENCES vehicle(id)
);
