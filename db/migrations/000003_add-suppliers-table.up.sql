CREATE TABLE "suppliers" (
	"id" uuid PRIMARY KEY,
	"name" text NOT NULL,
	"debt" numeric DEFAULT 0,
	"created_at" timestamptz DEFAULT (now()),
	"updated_at" timestamptz DEFAULT (now()) 
);