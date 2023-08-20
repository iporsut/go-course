CREATE TABLE IF NOT EXISTS todos (
   id UUID PRIMARY KEY,
   title VARCHAR (200) NOT NULL,
   created_at timestamp without time zone,
   completed_at timestamp without time zone
);
