-- Create "cabinets" table
CREATE TABLE "cabinets" (
  "id" bigserial NOT NULL,
  "uuid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_cabinets_uuid" to table: "cabinets"
CREATE INDEX "idx_cabinets_uuid" ON "cabinets" ("uuid");
