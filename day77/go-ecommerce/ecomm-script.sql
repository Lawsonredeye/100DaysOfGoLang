-- Database schema for go-ecommerce API

CREATE DATABASE IF NOT EXISTS go-ecommerce-db;

-- switch to the go-ecommerce-db
USE go-ecommerce-db;

-- Users information schema
CREATE TABLE IF NOT EXISTS users (
	id int auto_increment,
	username varchar(256) not null unique,
	email varchar(256) not null unique,
	password_hash varchar(256) not null,
	created_at datetime(3),
	updated_at datetime(3),
	role varchar(10),
	primary key(id)
);

-- Category Schema
CREATE TABLE IF NOT EXISTS categories (
	id int auto_increment,
	name varchar(256),

	primary key(id)
);

-- Product schema
CREATE TABLE IF NOT EXISTS products (
	id int auto_increment,
	name varchar(256),
	description text(1000),
	sku varchar(100),
	category_id int,
	color varchar(50),
	product_size int,
	price float,
	quantity int,
	created_at datetime,
	updated_at datetime,

	primary key(id),
	foreign key(category_id) references categories(id), 
);

-- creating table for orders
CREATE TABLE IF NOT EXISTS orders (
	id int auto_increment,
	user_id int,
	total_amount int,
	status string,
	created_at datetime,
	updated_at dateime,

	primary key(id),
	foreign key(category_id) references category(id), 
);

CREATE TABLE IF NOT EXISTS order_items (
	id int auto_increment,
	order_id int,
	product_id int,
	quantity int,
	price int,

	primary key(id),
	foreign key(order_id) references orders(id),
	foreign key(product_id) references products(id)
);
