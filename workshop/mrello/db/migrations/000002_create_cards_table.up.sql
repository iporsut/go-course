CREATE TABLE IF NOT EXISTS cards (
        id UUID PRIMARY KEY,
        title TEXT UNIQUE NOT NULL,
        description TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by_user_id UUID NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        updated_by_user_id UUID NOT NULL,
        CONSTRAINT fk_created_by_user_id
            FOREIGN KEY(created_by_user_id)
            REFERENCES users(user_id),
        CONSTRAINT fk_updated_by_user_id
            FOREIGN KEY(updated_by_user_id)
            REFERENCES users(user_id)
);
