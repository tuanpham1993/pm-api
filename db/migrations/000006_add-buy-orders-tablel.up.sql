CREATE TABLE "buy_orders" (
	"id" uuid PRIMARY KEY,
	"supplier_id" uuid,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);

ALTER TABLE "buy_orders" ADD FOREIGN KEY ("supplier_id") REFERENCES "suppliers" ("id") ON DELETE SET NULL;

CREATE TABLE "buy_order_items" (
	"id" uuid PRIMARY KEY,
	"buy_order_id" uuid NOT NULL,
    "product_id" uuid NOT NULL,
    "quantity" numeric NOT NULL,
	"price" numeric NOT NULL,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);

ALTER TABLE "buy_order_items" ADD FOREIGN KEY ("buy_order_id") REFERENCES "buy_orders" ("id") ON DELETE SET NULL;
ALTER TABLE "buy_order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE SET NULL;