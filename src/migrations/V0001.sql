create table if not exists person (
    id UUID PRIMARY KEY,
    name varchar(50) not null,
    cpf varchar(11) not null,
    email varchar(100),
    created_at timestamp not null,
    updated_at timestamp not null
);