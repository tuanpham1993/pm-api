CREATE TABLE "payments" (
	"id" uuid PRIMARY KEY,
	"supplier_id" uuid,
    "amount" numeric,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);

ALTER TABLE "payments" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id") ON DELETE SET NULL;
