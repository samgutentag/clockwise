-- +migrate Up
create table users (
    handle text primary key,
    tx_hash text not null
);

-- create table tracks (
--     id text primary key,
--     title text not null,
--     stream_url text not null,
--     description text not null,
--     user_id text not null,
--     tx_hash text not null,
--     foreign key (user_id) references users (id) on delete cascade
-- );

-- +migrate Down
drop table if exists users;

drop table if exists tracks;
