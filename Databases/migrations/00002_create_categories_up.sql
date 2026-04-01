create table if not exists categories(
    id serial primary key,
    name varchar(255) not null ,
    user_id int references users(id) on delete cascade,
    created_at timestamp default now(),
    unique(user_id,name)
);

create index if not exists idx_user_categories on categories(user_id);

