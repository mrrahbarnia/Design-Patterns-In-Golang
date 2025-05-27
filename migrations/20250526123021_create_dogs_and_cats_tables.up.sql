CREATE TABLE breeders (
    id SERIAL PRIMARY KEY,
    breeder_name TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    prov_state TEXT NOT NULL,
    country TEXT NOT NULL,
    zip_postal_code TEXT NOT NULL,
    phone TEXT,
    email TEXT,
    active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE dog_breeds (
    id SERIAL PRIMARY KEY,
    breed TEXT NOT NULL,
    weight_low_lbs INTEGER,
    weight_high_lbs INTEGER,
    average_weight INTEGER,
    lifespan INTEGER,
    description TEXT,
    alternate_names TEXT,
    geographic_origin TEXT
);

CREATE TABLE cat_breeds (
    id SERIAL PRIMARY KEY,
    breed TEXT NOT NULL,
    weight_low_lbs INTEGER,
    weight_high_lbs INTEGER,
    average_weight INTEGER,
    lifespan INTEGER,
    description TEXT,
    alternate_names TEXT,
    geographic_origin TEXT
);

CREATE TABLE dogs (
    id SERIAL PRIMARY KEY,
    breed_id INTEGER NOT NULL REFERENCES dog_breeds(id) ON DELETE CASCADE,
    dog_name TEXT NOT NULL,
    color TEXT,
    breeder_id INTEGER NOT NULL REFERENCES breeders(id) ON DELETE CASCADE,
    date_of_birth DATE,
    spayed_or_neutered BOOLEAN DEFAULT false,
    description TEXT,
    weight INTEGER
);

CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    breed_id INTEGER NOT NULL REFERENCES cat_breeds(id) ON DELETE CASCADE,
    cat_name TEXT NOT NULL,
    color TEXT,
    breeder_id INTEGER NOT NULL REFERENCES breeders(id) ON DELETE CASCADE,
    date_of_birth DATE,
    spayed_or_neutered BOOLEAN DEFAULT false,
    description TEXT,
    weight INTEGER
);

-- This table seems more like a generic read-only structure, not for relational mapping.
-- We'll create it anyway, without constraints.
CREATE TABLE pets (
    species TEXT,
    breed TEXT,
    min_weight INTEGER,
    max_weight INTEGER,
    description TEXT,
    lifespan INTEGER
);
