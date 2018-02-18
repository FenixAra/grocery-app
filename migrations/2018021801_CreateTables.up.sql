CREATE TABLE category (
	id UUID PRIMARY KEY,
	name TEXT,
	description TEXT,
	parent_id TEXT
);

CREATE INDEX category_parent_id_idx ON category (parent_id);

CREATE TABLE user (
	id UUID PRIMARY KEY,
	name TEXT,
	email TEXT,
	mobile TEXT,
	type TEXT,
	age INTEGER
);

CREATE INDEX user_age_idx ON user (age);


CREATE TABLE price_card (
	id UUID PRIMARY KEY,
	name TEXT,
	code TEXT,
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
	ixclusion   TEXT[]
);

CREATE INDEX discount_inclusion_idx ON discount USING GIN (inclusion);
CREATE INDEX discount_exclusion_idx ON discount USING GIN (exclusion);

CREATE TABLE item (
	id          UUID PRIMARY KEY,
	name        TEXT,
	code        TEXT,
	description TEXT,
	price_card_id UUID REFERENCES price_card (id),
	tags        []TEXT
);

CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE tag (
	id UUID PRIMARY KEY,
	name TEXT,
	code TEXT
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
	user_id UUID REFERENCES user (id),
	status TEXT
);

CREATE INDEX register_status_idx ON register (status);

CREATE TABLE order (
	id           UUID PRIMARY KEY,
	user_id      UUID REFERENCES user (id),
	register_id  UUID REFERENCES register (id),
	employee_id  UUID REFERENCES user (id),
	amount       INTEGER,
	bill 		 JSON,
	inventories  []TEXT,
	created_at   TIMESTAMP WITHOUT TIMEZONE
);
