CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 

    name TEXT NOT NULL CHECK (char_length(trim(name)) BETWEEN 3 AND 100), 

    email TEXT NOT NULL, 

    department TEXT NOT NULL, 
    
    password_hash TEXT NOT NULL, 

    role TEXT NOT NULL DEFAULT 'requester' CHECK (role IN ('requester', 'technician', 'admin')),

    is_active BOOLEAN NOT NULL DEFAULT TRUE, 

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(), 

    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()   
);

CREATE UNIQUE INDEX user_email_lower_unique_idx ON users (lower(email));