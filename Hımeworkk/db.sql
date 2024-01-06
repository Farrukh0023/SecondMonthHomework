create table users(
    id uuid primary key not null ,
    first_name varchar(30) ,
    last_name  varchar(30) ,
    email varchar not null ,
    phone varchar(13) not null
);
create table orders(
    id uuid primary key not null,
    amount int default 0,
    user_id uuid references users(id) ON DELETE CASCADE,
    created_at timestamp default current_timestamp

);
create table products(
    id uuid primary key not null ,
    name varchar(50),
    price int
);
create table order_products(
    id uuid primary key not null,
    order_id uuid references orders(id) ON DELETE CASCADE,
    product_id uuid references products(id) on delete cascade ,
    quantity int default 0,
    price int
);