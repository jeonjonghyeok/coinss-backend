CREATE TABLE public.users (
    uid serial PRIMARY KEY,
    email varchar (40) UNIQUE,
    password varchar (80),
    name varchar (40),
    favorites varchar (200),
    access_key varchar (50),
    secret_key varchar (50)
);


CREATE TABLE public.search (
    sid serial PRIMARY KEY,
    search varchar (200),
    count int
);

CREATE TABLE public.coins (
    cid serial PRIMARY KEY,
    name varchar(100),
)