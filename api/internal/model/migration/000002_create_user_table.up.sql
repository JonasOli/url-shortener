create table users ( 
    id serial primary key,
    name text not null,
    created_at timestamp default now()
);

alter table urls
add column created_by int references users not null;
