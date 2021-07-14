CREATE TABLE public.users (
    uid serial PRIMARY KEY,
    email varchar (40) UNIQUE,
    password varchar (80),
    name varchar (40),
    favorites varchar (200),
    access_key varchar (50),
    secret_key varchar (50)
);