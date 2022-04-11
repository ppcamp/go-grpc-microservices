CREATE TABLE IF NOT EXISTS "password" (
	id serial4 NOT NULL,
	user_id varchar NOT NULL,
	user_password varchar(200) NOT NULL,
	CONSTRAINT password_pk PRIMARY KEY (id)
);