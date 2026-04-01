create table
    if not exists expenses (
        id serial primary key,
        amount decimal(10,2) not null,
        merchant varchar(255) not null,
        category_id int references categories (id) on delete cascade not null,
        description varchar(255) not null,
        user_id int references users (id) on delete cascade not null,
        date timestamp default now(),
        source varchar(255) not null,
        created_at timestamp default now() not null
    );
    create index if not exists idx_user_expenses on expenses(user_id);