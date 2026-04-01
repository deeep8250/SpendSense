create table if not exists users(
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) not null unique,
    hashed_password varchar(255) not null,
    created_at timestamp default now()
);
