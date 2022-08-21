CREATE TABLE user_session (
    id serial PRIMARY KEY,
    user_id bigint references app_user(id),
    token varchar(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);