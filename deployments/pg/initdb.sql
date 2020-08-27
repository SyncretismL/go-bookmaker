CREATE TABLE public.sports (
    id bigserial PRIMARY KEY,
    sport text UNIQUE NOT NULL,
    coefficient numeric(5,2) NOT NULL
);