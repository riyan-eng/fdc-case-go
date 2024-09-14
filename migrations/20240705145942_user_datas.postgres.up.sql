create table user_datas(
    id serial not null,
    uuid uuid not null default gen_random_uuid(),
    user_uuid uuid,
    role_code varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    
    primary key(id),
    unique(uuid),
    unique(user_uuid, role_code),
    constraint fk_user foreign key(user_uuid) references users(uuid) on delete cascade,
    constraint fk_role foreign key(role_code) references roles(code) on delete no action
);