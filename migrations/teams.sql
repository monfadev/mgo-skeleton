DROP TYPE IF EXISTS role_team;

CREATE TYPE role_team AS ENUM ('admin', 'operator');

CREATE TABLE public.teams (
  id SERIAL PRIMARY KEY,  -- Auto-incrementing integer for ID
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,  -- Unique constraint for email
  role role_team NOT NULL,
  user_id INTEGER not null,
  updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL  
);