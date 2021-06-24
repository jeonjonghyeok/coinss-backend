CREATE TABLE public.user (
    uid serial PRIMARY KEY,
    email varchar (40) UNIQUE,
    password varchar (40),
    name varchar (40),
    phone_number varchar (40),
    access_key varchar (50),
    secret_key varchar (50)
);