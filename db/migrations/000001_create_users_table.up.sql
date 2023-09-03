CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES auth.users (id)
    
);