CREATE DATABASE db_warung_makan;

-- SELECT pg_terminate_backend(pg_stat_activity.pid)
-- FROM pg_stat_activity
-- WHERE pg_stat_activity.datname = 'db_warung_makan';
-- DROP DATABASE db_warung_makan;

CREATE TABLE m_login(
	id varchar(50) PRIMARY KEY NOT NULL,
	username varchar(50),
	password varchar(50),
	id_customer varchar(50)
);

CREATE TABLE m_customer(
	id varchar(50) PRIMARY KEY NOT NULL,
	customer_name VARCHAR(50),
	email varchar(50),
	active_member boolean,
	address varchar(50)
);

CREATE TABLE m_order(
	id varchar(250) PRIMARY KEY NOT NULL,
	customer_id VARCHAR(50),
	table_id integer,
	paid_status boolean,
	total_price integer,
	order_detail_id integer,
	reservation_number integer NULL
);

CREATE TABLE m_suplier(
	id serial PRIMARY KEY NOT NULL,
	suplier_name VARCHAR(50),
	faktur_number varchar(50),
	receive_date varchar(50),
	id_food_stuff varchar(50),
	price_total integer,
	employee_receiver integer
);

CREATE TABLE m_suplier_detail(
	id serial PRIMARY KEY NOT NULL,
	quantity integer,
	price_total integer
);

CREATE TABLE m_employee(
	id serial PRIMARY KEY NOT NULL,
	email VARCHAR(50),
	address varchar(50),
	phone_number varchar(50),
	role_employee varchar(50),
	employee_name varchar(50)
);

CREATE TABLE m_table(
	id serial PRIMARY KEY NOT NULL,
	table_number varchar(100),
	reservation_status integer
);

CREATE TABLE m_order_detail(
	id serial PRIMARY KEY NOT NULL,
	menu_id INTEGER,
	quantity integer,
	order_id integer
);

CREATE TABLE m_reservation(
	id serial PRIMARY KEY NOT NULL,
	reservation_date date,
	table_number integer,
	customer_id VARCHAR(50),
);

CREATE TABLE m_menu(
	id serial PRIMARY KEY NOT NULL,
	menu_name varchar(50),
	price integer
);


