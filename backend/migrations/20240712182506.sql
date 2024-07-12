-- Create "cabinets" table
CREATE TABLE "cabinets" (
  "id" bigserial NOT NULL,
  "uuid" uuid NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_cabinets_uuid" to table: "cabinets"
CREATE INDEX "idx_cabinets_uuid" ON "cabinets" ("uuid");
