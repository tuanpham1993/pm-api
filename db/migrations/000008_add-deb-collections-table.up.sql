CREATE TABLE "debt_collections" (
	"id" uuid PRIMARY KEY,
	"customer_id" uuid,
    "amount" numeric,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);

ALTER TABLE "debt_collections" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") ON DELETE SET NULL;
