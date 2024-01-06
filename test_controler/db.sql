create table cars (
    id uuid primary key not null ,
    model varchar(30),
    brand varchar(30),
    year int not null
);
create table driver(
    id uuid primary key not null ,
    full_name varchar not null ,
    phone varchar(13) not null ,
    car_id uuid references cars(id) ON DELETE CASCADE
);