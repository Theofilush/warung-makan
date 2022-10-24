INSERT INTO public.m_customer(
	id, customer_name, email, active_member, address)
	VALUES ('1', 'Vicky', 'email@gmail.com', TRUE, 'Jakarta'),
	('2', 'Kenny', 'email@gmail.com', TRUE, 'Jakarta'),
	('3', 'Karl', 'email@gmail.com', TRUE, 'Jakarta'),
	('4', 'Bagas', 'email@gmail.com', TRUE, 'Jakarta'),
	('5', 'Zainal', 'email@gmail.com', TRUE, 'Jakarta')
	;
	
INSERT INTO public.m_login(
	id, username, password, id_customer)
	VALUES ('1', 'vicky', 'vicky', '1'),
	('2', 'kenny', 'kenny', '2'),
	('3', 'karl', 'karl', '3'),
	('4', 'bagas', 'bagas', '4'),
	('5', 'zainal', 'zainal', '5')
	;	
	
INSERT INTO public.m_employee(
	id, email, address, phone_number, role_employee, employee_name)
	VALUES (1, 'email@gmail.com', 'Jakarta', '089765432110', 'Koki', 'Jason'),
	(2, 'email@gmail.com', 'Bandung', '089765432110', 'Manager', 'Lukas'),
	(3, 'email@gmail.com', 'Bogor', '089765432110', 'Waiter', 'Agra'),
	(4, 'email@gmail.com', 'Bandung', '089765432110', 'Waitress', 'Anisa'),
	(5, 'email@gmail.com', 'Bogor', '089765432110', 'Kasir', 'Aden'),
	(6, 'email@gmail.com', 'jakarta', '089765432110', 'Kasir', 'Cairo'),
	(7, 'email@gmail.com', 'Bekasi', '089765432110', 'Koki', 'Edison'),
	(8, 'email@gmail.com', 'Jakarta', '089765432110', 'Waiter', 'Reynald')
	;

INSERT INTO public.m_menu(
	id, menu_name, price)
	VALUES (1, 'Bistik Daging Sapi', 100000),
	(2, 'Bistik Galantine', 100000),
	(3, 'Bistik Udang Spesial', 100000),
	(4, 'Bistik  Iga Spesial', 100000),
	(5, 'Bistik Iga', 100000),
	(6, 'Bistik Lidah', 100000),
	(7, 'Bistik Ayam Spesial', 100000),
	(8, 'Bistik Ikan', 100000),
	(9, 'Nasi Goreng Bistik Labuan Bajo', 100000)
	;

INSERT INTO public.m_order(
	id, customer_id, table_id, paid_status, total_price, order_detail_id, reservation_number)
	VALUES ('1', '1', '1', TRUE, 100000, 1, NULL),
	('2', '2', '2', TRUE, 100000, 2, NULL),
	('3', '3', '3', TRUE, 100000, 3, NULL),
	('4', '4', '4', TRUE, 100000, 4, NULL),
	('5', '5', '5', TRUE, 100000, 5, NULL),
	('6', '6', '6', TRUE, 100000, 6, NULL),
	('7', '7', '7', TRUE, 100000, 7, NULL)
	;
	
INSERT INTO public.m_order_detail(
	id, menu_id, quantity,order_id)
	VALUES (1, 1, 1, 1),
	(2, 2, 2, 2),
	(3, 3, 3, 3),
	(4, 4, 4, 4),
	(5, 5, 5, 5),
	(6, 6, 6, 6),
	(7, 7, 7, 7);
	
INSERT INTO public.m_reservation(
	id, reservation_date, table_number, customer_id)
	VALUES (1, '2022-10-10', 1, '1'),
	 (2, '2022-10-11', 2, '2'),
	 (3, '2022-10-12', 3, '3'),
	 (4, '2022-10-13', 4, '4'),
	 (5, '2022-10-14', 5, '5')
	 ;

INSERT INTO public.m_suplier(
	id, suplier_name, faktur_number, receive_date, id_food_stuff, price_total, employee_receiver)
	VALUES (1, 'PT Food Segar', '1234567890', '2022-10-10', 1, 120000, 1),
	(2, 'PT Food Segar', '1234567891', '2022-10-11', 2, 120000, 2),
	(3, 'PT Food Segar', '1234567892', '2022-10-12', 3, 120000, 3);


INSERT INTO public.m_suplier_detail(
	id, quantity, price_total)
	VALUES (1, 10, 120000),
	(2, 10, 120000),
	(3, 10, 120000);
	
INSERT INTO public.m_table(
	id, table_number, reservation_status)
	VALUES (1, '1', 1),
	(2, '2', 1),
	(3, '3', 1),
	(4, '4', 1),
	(5, '5', 1),
	(6, '6', 1),
	(7, '7', 1),
	(8, '8', 1),
	(9, '9', 1),
	(10, '10', 1),
	(11, '11', 1),
	(12, '12', 1),
	(13, '13', 1),
	(14, '1', 1),
	(15, '15', 1);
	
	
SELECT * FROM m_customer;
SELECT * FROM m_employee;
SELECT * FROM m_login;
SELECT * FROM m_menu;
SELECT * FROM m_order;
SELECT * FROM m_order_detail;
SELECT * FROM m_reservation;
SELECT * FROM m_suplier;
SELECT * FROM m_suplier_detail;
SELECT * FROM m_table;






select * from m_order join m_order_detail on m_order.order_detail_id = m_order_detail.order_id



select 
m_order.id, 
m_order.customer_id, 
m_order.table_id, 
m_order.paid_status, 
m_order.total_price, 
m_order.order_detail_id, 
m_order_detail.menu_id, 
m_order_detail.quantity 
from m_order 
join m_order_detail on m_order.order_detail_id = m_order_detail.order_id
where m_order.id='1234567892'



