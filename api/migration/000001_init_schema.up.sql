create table urls (
    id serial primary key,
    original text not null,
    short_code text unique not null,
    visit_count int default 0,
    created_at timestamp default now()
);
