CREATE TABLE IF NOT EXISTS "passwords" (
	id BIGSERIAL NOT NULL,
	user_id VARCHAR(200) NOT NULL,
	user_password VARCHAR(200) NOT NULL,
	active BOOLEAN NOT NULL DEFAULT false,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NULL,
	CONSTRAINT passwords_pk PRIMARY KEY (id)
);
CREATE INDEX passwords_user_id_idx ON passwords USING btree (user_id);