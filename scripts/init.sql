CREATE UNLOGGED TABLE "clients" (
	"id" SERIAL PRIMARY KEY NOT NULL,
	"limit" integer NOT NULL,
	"balance" integer DEFAULT 0 NOT NULL
  CONSTRAINT can_make CHECK (clients.balance >= -clients.limit)
);

CREATE UNLOGGED TABLE "transactions" (
	"id" SERIAL PRIMARY KEY NOT NULL,
	"value" integer NOT NULL,
	"description" text NOT NULL,
	"client_id" integer NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients(id)
);

-- TODO: Indexes

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