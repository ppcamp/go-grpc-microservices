CREATE TABLE IF NOT EXISTS "password" (
	id BIGSERIAL NOT NULL,
	user_id VARCHAR(200) NOT NULL,
	user_password VARCHAR(200) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NULL,
	CONSTRAINT password_pk PRIMARY KEY (id)
);
CREATE INDEX password_user_id_idx ON public.password USING btree (user_id);