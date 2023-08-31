CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="America/Sao_Paulo";

-- Create  table
CREATE TABLE public.vagas (
    id SERIAL PRIMARY KEY,
    porte VARCHAR(10) NOT null,
    ocupacao TIMESTAMP DEFAULT NOW(),
    dono VARCHAR(255),
    telefone VARCHAR(255),
    tipo_de_veiculo VARCHAR(255),
    placa VARCHAR(10),
    modelo VARCHAR(255)
);
