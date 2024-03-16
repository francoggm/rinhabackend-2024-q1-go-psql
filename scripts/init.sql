CREATE UNLOGGED TABLE "clients" (
	"id" SERIAL PRIMARY KEY NOT NULL,
	"limit" integer NOT NULL,
	"balance" integer DEFAULT 0 NOT NULL
  CONSTRAINT make_transaction CHECK (clients.balance >= -clients.limit)
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

-- TODO: Indexes

CREATE OR REPLACE FUNCTION make_transaction(
  client_id INTEGER,
  transaction_value NUMERIC,
	transaction_description TEXT,
	transaction_type CHAR
) RETURNS TABLE (new_balance NUMERIC, new_limit NUMERIC) AS
$$
DECLARE
    client_balance NUMERIC;
    client_limit NUMERIC;
		debito_value NUMERIC;
BEGIN
	SELECT clients.balance, clients.limit INTO client_balance, client_limit
	FROM clients
	WHERE id = client_id;

	IF NOT FOUND THEN
		RAISE EXCEPTION 'client not found';
	END IF;

	IF transaction_type = 'd' THEN
		client_balance := client_balance - transaction_value;
	ELSE
		client_balance := client_balance + transaction_value;
	END IF;

	UPDATE clients
	SET balance = client_balance
	WHERE id = client_id;

	INSERT INTO transactions (client_id, value, description, type)
  VALUES (client_id, transaction_value, transaction_description, transaction_type);

	RETURN QUERY SELECT client_balance, client_limit;
END;
$$
LANGUAGE plpgsql;

DO $$
BEGIN
	INSERT INTO "clients" ("id", "limit")
	VALUES
		(1, 100000),
    (2, 80000),
    (3, 1000000),
    (4, 10000000),
    (5, 500000);
END;
$$;