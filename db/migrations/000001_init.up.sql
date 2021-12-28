CREATE TABLE "products" (
	"id" uuid PRIMARY KEY,
	"name" text NOT NULL,
	"unit" text NOT NULL,
	"quantity" numeric DEFAULT 0,
	"input_price" numeric,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);