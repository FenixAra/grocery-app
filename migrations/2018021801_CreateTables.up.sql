CREATE TABLE category (
	id UUID PRIMARY KEY,
	name TEXT,
	description TEXT,
	parent_id TEXT
);

CREATE INDEX category_parent_id_idx ON category (parent_id);

CREATE TABLE account (
	id UUID PRIMARY KEY,
	name TEXT,
	email TEXT,
	mobile TEXT,
	type TEXT,
	age INTEGER
);

CREATE INDEX account_age_idx ON account (age);
CREATE UNIQUE INDEX account_email_idx ON account (email);
CREATE UNIQUE INDEX account_mobile_idx ON account (mobile);


CREATE TABLE price_card (
	id UUID PRIMARY KEY,
	code TEXT,
	name TEXT,
	description TEXT,
	total INTEGER
);


CREATE TABLE line_item (
	id UUID PRIMARY KEY,
	price_card_id UUID REFERENCES price_card (id),
	name TEXT,
	code TEXT,
	description TEXT,
	amount INTEGER
);

CREATE INDEX line_item_price_card_id_idx ON line_item (price_card_id);

CREATE TABLE discount(
	id          UUID PRIMARY KEY,
	name        TEXT,
	code        TEXT,
	description TEXT,
	type        TEXT,
	amount      INTEGER,
	percent     INTEGER,
	inclusion   TEXT[],
	exclusion   TEXT[]
);

CREATE INDEX discount_inclusion_idx ON discount USING GIN (inclusion);
CREATE INDEX discount_exclusion_idx ON discount USING GIN (exclusion);

CREATE TABLE item (
	id          UUID PRIMARY KEY,
	name        TEXT,
	code        TEXT,
	description TEXT,
	price_card_id UUID REFERENCES price_card (id),
	tags        TEXT[]
);

CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE tag (
	id UUID PRIMARY KEY,
	name TEXT,
	type TEXT
);

CREATE INDEX tag_name_idx ON tag USING GIST (name gist_trgm_ops);

CREATE TABLE inventory (
	bar_code    UUID PRIMARY KEY,
	item_id     UUID REFERENCES item (id),
	name        TEXT,
	descritpion TEXT,
	status      TEXT
);

CREATE INDEX inventory_status_idx ON inventory (status);

CREATE TABLE register (
	id UUID PRIMARY KEY,
	account_id UUID REFERENCES account (id),
	status TEXT
);

CREATE INDEX register_status_idx ON register (status);

CREATE TABLE booking (
	id           UUID PRIMARY KEY,
	account_id   UUID REFERENCES account (id),
	register_id  UUID REFERENCES register (id),
	employee_id  UUID REFERENCES account (id),
	amount       INTEGER,
	bill 		 JSON,
	inventories  TEXT[],
	created_at   TIMESTAMP WITHOUT TIME ZONE
);


CREATE TABLE payment (
	id UUID PRIMARY KEY,
	booking_id UUID REFERENCES booking(id),
	mode TEXT,
	payment_ref TEXT,
	amount INTEGER,
	status TEXT
);
