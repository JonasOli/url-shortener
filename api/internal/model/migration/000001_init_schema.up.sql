create table users ( 
    id serial primary key,
    name text not null,
    email text not null unique,
    password text not null,
    salt text not null,
    created_at timestamp default now()
);

create table urls (
    id serial primary key,
    original text not null,
    short_code text unique not null,
    visit_count int default 0,
    created_at timestamp default now(),
    created_by int references users not null
);