drop database if exists auth;

create database auth;

use auth;

create table users (
	username varchar(32) not null,
	password varchar(128) not null,
	primary key(username)
);

insert into users(username, password) values
	('fred', sha1('fred')),
	('wilma', sha1('wilma')),
	('barney', sha1('barney')),
	('betty', sha1('betty'));

