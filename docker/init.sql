CREATE SCHEMA cart;

ALTER SCHEMA cart OWNER TO postgres;

CREATE TABLE cart.cart (
    id integer NOT NULL
);
ALTER TABLE ONLY cart.cart
    ADD CONSTRAINT cart_pk PRIMARY KEY (id);
	
ALTER TABLE cart.cart OWNER TO postgres;

CREATE SEQUENCE cart.cart_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE cart.cart_id_seq OWNER TO postgres;
ALTER SEQUENCE cart.cart_id_seq OWNED BY cart.cart.id;

ALTER TABLE ONLY cart.cart ALTER COLUMN id SET DEFAULT nextval('cart.cart_id_seq'::regclass);

CREATE TABLE cart.cart_item (
    id integer NOT NULL,
    product character varying(255) NOT NULL CHECK (product <> ''),
    quantity integer NOT NULL CHECK (quantity >= 0),
    cart_id integer NOT NULL
);

ALTER TABLE cart.cart_item OWNER TO postgres;

CREATE SEQUENCE cart.cart_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE cart.cart_item_id_seq OWNER TO postgres;
ALTER SEQUENCE cart.cart_item_id_seq OWNED BY cart.cart_item.id;

ALTER TABLE ONLY cart.cart_item ALTER COLUMN id SET DEFAULT nextval('cart.cart_item_id_seq'::regclass);

ALTER TABLE ONLY cart.cart_item
    ADD CONSTRAINT cart_item_pk PRIMARY KEY (id);

ALTER TABLE ONLY cart.cart_item
    ADD CONSTRAINT cart_item_fk FOREIGN KEY (cart_id) REFERENCES cart.cart(id);
