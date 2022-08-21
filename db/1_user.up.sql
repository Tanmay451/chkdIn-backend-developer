CREATE TABLE app_user (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
	salt varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_active boolean DEFAULT TRUE
);