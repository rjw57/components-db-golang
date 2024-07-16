--liquibase formatted sql

--changeset rjw57:1
CREATE TABLE "public"."cabinets" (
  "id" bigserial NOT NULL,
  "uuid" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT current_timestamp,
  "updated_at" timestamptz NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY ("id")
);

CREATE INDEX "idx_cabinets_uuid" ON "public"."cabinets" ("uuid");

--changeset rjw57:2
CREATE FUNCTION update_updated_at()
RETURNS TRIGGER
LANGUAGE plpgsql
AS '
BEGIN
    NEW.updated_at = current_timestamp;
    RETURN NEW;
END;';

CREATE TRIGGER update_updated_at
    BEFORE UPDATE
    ON
        cabinets
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();
