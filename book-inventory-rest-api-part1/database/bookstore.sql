drop database bookstore;
create database bookstore;
use bookstore;
create table books(
	ID varchar(250) autoincrement,
    Title varchar (250),
    Author varchar(250)
);

insert into books (ID, Title,Author) values ('1', 'atomic habits','napolean hill');
insert into books (ID, Title,Author) values ('2', 'how to influence friend and win people','dale carnegie');
select * from books