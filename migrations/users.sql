CREATE TABLE public.users (
  id SERIAL PRIMARY KEY,  -- Auto-incrementing integer for ID
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,  -- Unique constraint for email
  role enum_role NOT NULL,
  updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TYPE enum_role AS ENUM ('superadmin','admin', 'operator');