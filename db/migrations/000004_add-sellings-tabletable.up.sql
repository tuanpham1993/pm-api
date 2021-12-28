CREATE TABLE "sell_orders" (
	"id" uuid PRIMARY KEY,
	"customer_id" uuid,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);

ALTER TABLE "sell_orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") ON DELETE SET NULL;