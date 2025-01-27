CREATE TYPE type AS ENUM ('military vehicle', 'weapon');

CREATE TABLE techniques (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    model VARCHAR(50) NOT NULL UNIQUE,
    type type NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE bullets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    caliber FLOAT NOT NULL UNIQUE,
    type type NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TYPE fuel AS ENUM ('petrol', 'diesel');
CREATE TABLE fuels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type fuel NOT NULL UNIQUE,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);
