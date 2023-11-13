create table if not exists users
(
    id      bigint primary key generated always as identity,
    login   varchar(200) not null unique,
    hashed_password varchar(200) not null
);
insert into users (login, hashed_password)
values ('12', 'c20ad4d76fe97759aa27a0c99bff6710');

select * from users

drop table logs
create table if not exists logs
(
    id      bigint primary key generated always as identity,
    device_name   varchar(200) not null,
	port integer,
	port_state varchar(200) not null,
	bool_change bool,
    date_time timestamp not null
);

select * from logs