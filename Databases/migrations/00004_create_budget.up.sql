create table
    if not exists budgets (
        id serial primary key,
        user_id int references users (id) on delete cascade not null,
        category_id int references categories (id) on delete cascade not null,
        amount decimal(10,2) not null,
        month int not null,
        year int not null,
        created_at timestamp default now (),
        unique (user_id, category_id,month,year)
    );

create index if not exists idx_user_budget on budgets (user_id);