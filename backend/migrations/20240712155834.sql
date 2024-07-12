-- Create "cabinets" table
CREATE TABLE "cabinets" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_cabinets_deleted_at" to table: "cabinets"
CREATE INDEX "idx_cabinets_deleted_at" ON "cabinets" ("deleted_at");
