CREATE TABLE ticket (
    Id uuid primary key not null ,
    where_from varchar(30) not null ,
    where_to varchar(30) not null ,
    date varchar(30) not null
);

CREATE TABLE users (
    id uuid primary key not null,
    first_name varchar(30) default null ,
    last_name varchar(30)  default null,
    email varchar(40) not null,
    ticket_id uuid references ticket(Id)
);

select u.first_name, u.last_name, u.email, t.from, t.to, t.date from user as u join ticket as t on u.ticket_id = t.Id
