CREATE EXTENSION IF NOT EXISTS "uuid-ossp"
drop table users cascade;
drop table locations cascade;
drop table events cascade;
drop table categories cascade;

create table users(
id uuid primary key DEFAULT uuid_generate_v4(),
name text,
email text,
password text,
createdAt TIMESTAMP default CURRENT_TIMESTAMP ,
updatedAt TIMESTAMP default CURRENT_TIMESTAMP 
);

create table categories(
id uuid primary key DEFAULT uuid_generate_v4(),
name text,
createdAt TIMESTAMP default CURRENT_TIMESTAMP ,
updatedAt TIMESTAMP default CURRENT_TIMESTAMP 
);

create table locations(
id uuid primary key DEFAULT uuid_generate_v4(),
name text,
addres text,
createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create table events(
id uuid primary key DEFAULT uuid_generate_v4(),
title text,
description text,
date TIMESTAMP,
organizerId uuid references users(id) ON DELETE CASCADE ,
categoryId uuid references categories(id) ON DELETE SET NULL,
locationId uuid references locations(id) ON DELETE SET NULL,
createdAt TIMESTAMP default CURRENT_TIMESTAMP,
updatedAt TIMESTAMP default CURRENT_TIMESTAMP
);


