CREATE TABLE departments (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    department_id INTEGER NOT NULL REFERENCES departments(id),
    role VARCHAR(50) NOT NULL DEFAULT 'employee' CHECK (role IN ('employee', 'technician', 'admin')),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW() 
);

INSERT INTO departments (name)
VALUES
    ('Business Development'),
    ('Customer Service'),
    ('Finance & Accounting')
    ('Human Resources'), 
    ('IT'), 
    ('Legal & Compliance'), 
    ('Logistics'), 
    ('Management'), 
    ('Marketing'), 
    ('Operations'), 
    ('Procurement'), 
    ('Sales');
