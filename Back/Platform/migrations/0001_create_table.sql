CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="America/Sao_Paulo";

-- Create  table
CREATE TABLE public.clients (
    id UUID PRIMARY KEY,
    username varchar(255),
    name varchar(255),
    phone varchar(15),
    created_at timestamp DEFAULT current_timestamp
);

CREATE TABLE public.vehicle (
    id UUID PRIMARY KEY,
    owner varchar(255),
    title varchar(255),
    type varchar(255),
    plate varchar(10),
    created_at timestamp DEFAULT current_timestamp
);

CREATE TABLE public.space (
    id UUID PRIMARY KEY,
    vehicle_id UUID,
    type varchar(255),
    created_at timestamp DEFAULT current_timestamp,
    FOREIGN KEY (vehicle_id) REFERENCES vehicle(id)
);

CREATE TABLE public.history (
    id UUID PRIMARY KEY,
    amount float,
    vehicle_id UUID,
    entry timestamp,
    exit timestamp,
    type varchar(255),
    created_at timestamp DEFAULT current_timestamp,
    FOREIGN KEY (vehicle_id) REFERENCES vehicle(id)
);