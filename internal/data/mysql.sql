drop table user_users;
drop table sessions;
drop table profiles;
drop table genders;
drop table users;

create table users
(
    id         serial primary key,
    uuid       varchar(64)  not null unique,
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at datetime     not null
);

create table user_users
(
    user_id      integer references users (id),
    friend_id    integer references users (id),
    created_at   datetime not null,
    confirmed_at datetime

);

create table genders
(
    id    smallint unsigned primary key,
    value varchar(32) not null unique
);

insert into genders (id, value)
VALUES (1, 'male'),
       (2, 'female');


create table profiles
(
    id         serial primary key,
    uuid       varchar(64) not null unique,
    user_id    integer unique references users (id),
    first_name varchar(64),
    last_name  varchar(64),
    age        smallint unsigned,
    gender     smallint unsigned references genders (id),
    interests  text,
    city       varchar(64),
    created_at timestamp   not null
);

create table sessions
(
    id         serial primary key,
    uuid       varchar(64) not null unique,
    email      varchar(255),
    user_id    integer references users (id),
    created_at timestamp   not null
);