CREATE TABLE "sell_order_items" (
	"id" uuid PRIMARY KEY,
	"sell_order_id" uuid NOT NULL,
    "product_id" uuid NOT NULL,
    "quantity" numeric NOT NULL,
	"price" numeric NOT NULL,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);

ALTER TABLE "sell_order_items" ADD FOREIGN KEY ("sell_order_id") REFERENCES "sell_orders" ("id") ON DELETE SET NULL;
ALTER TABLE "sell_order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE SET NULL;