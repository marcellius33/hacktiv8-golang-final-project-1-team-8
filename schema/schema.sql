create table to_dos
(
    id           uuid         not null primary key,
    title        varchar(255) not null,
    description  varchar(255) not null,
    completed_at timestamp(6) with time zone,
    deleted_at   timestamp(6) with time zone,
    created_at   timestamp(6) with time zone,
    updated_at   timestamp(6) with time zone
);