create table users(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    username varchar(255) not null,
    "password" text,
    is_active boolean not null default true,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(username)
);