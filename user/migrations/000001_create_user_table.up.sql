CREATE TABLE IF NOT EXISTS "users" (
	id UUID NOT NULL DEFAULT generate_uuid_v4(),
	first_name VARCHAR(50) NOT NULL,
	middle_name VARCHAR(130),
	last_name VARCHAR(20) NOT NULL,
	nick VARCHAR(20) NOT NULL,
	birthdate DATE NOT NULL,
	email VARCHAR(100),
	active BOOLEAN NOT NULL DEFAULT false,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);
CREATE INDEX users_id_idx ON users USING btree (id);