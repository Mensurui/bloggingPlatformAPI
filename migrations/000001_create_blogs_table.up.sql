CREATE TABLE IF NOT EXISTS blog (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    content text NOT NULL,
    tag text[] NOT NULL,
    updated_at timestamp(0) with time zone NOT NULL
);

