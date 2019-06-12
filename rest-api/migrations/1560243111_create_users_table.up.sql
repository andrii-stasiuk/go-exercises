CREATE TABLE public.users (
    id bigserial,
    email text NOT NULL UNIQUE,
    password bytea,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);