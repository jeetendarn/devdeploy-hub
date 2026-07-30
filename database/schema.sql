CREATE TABLE users (

    id UUID PRIMARY KEY,

    full_name VARCHAR(120) NOT NULL,

    email VARCHAR(120) UNIQUE NOT NULL,

    password VARCHAR(255) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE projects (

    id UUID PRIMARY KEY,

    name VARCHAR(120) NOT NULL,

    description TEXT,

    environment VARCHAR(50),

    status VARCHAR(50),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE deployments (

    id UUID PRIMARY KEY,

    project_id UUID REFERENCES projects(id),

    version VARCHAR(30),

    status VARCHAR(50),

    deployed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

