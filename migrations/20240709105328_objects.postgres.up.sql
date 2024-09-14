create table if not exists objects (
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    "name" varchar(255),
    "owner" uuid,
    "size" int,
    content_type varchar(255),
    "path" varchar,
    "url" varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid)
);