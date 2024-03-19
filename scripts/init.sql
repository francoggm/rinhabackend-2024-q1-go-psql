CREATE UNLOGGED TABLE "clients" (
	"id" SERIAL PRIMARY KEY NOT NULL,
	"limite" integer NOT NULL,
	"balance" integer DEFAULT 0 NOT NULL
  CONSTRAINT can_mk CHECK (balance >= -limite)
);

CREATE UNLOGGED TABLE "transactions" (
	"id" SERIAL PRIMARY KEY NOT NULL,
	"value" integer NOT NULL,
	"type" char NOT NULL,
	"description" text NOT NULL,
	"client_id" integer NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE INDEX idx_client_id ON clients (id);
CREATE INDEX idx_transactions_client_id_created_at ON transactions (client_id, created_at DESC);

CREATE OR REPLACE FUNCTION make_transaction(
  client_id INTEGER,
  transaction_value INTEGER,
	transaction_description TEXT,
	transaction_type CHAR
) RETURNS TABLE (new_balance INTEGER, new_limit INTEGER) AS
$$
DECLARE
    correct_value INTEGER;
BEGIN
	IF transaction_type = 'd' 
		THEN correct_value := -transaction_value;
		ELSE correct_value := transaction_value;
	END IF; 

	UPDATE clients
	SET balance = balance + correct_value
	WHERE id = client_id
	RETURNING balance, limite INTO new_balance, new_limit;

	INSERT INTO transactions (client_id, value, description, type)
  VALUES (client_id, transaction_value, transaction_description, transaction_type);
	
	RETURN QUERY SELECT new_balance, new_limit;
END;
$$
LANGUAGE plpgsql;

DO $$
BEGIN
	INSERT INTO "clients" ("id", "limite", "balance")
	VALUES
		(1, 100000, 0),
    (2, 80000, 0),
    (3, 1000000, 0),
    (4, 10000000, 0),
    (5, 500000, 0);
END;
$$;