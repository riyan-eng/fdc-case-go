create table roles(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    code varchar(255),
    "name" varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(code)
);