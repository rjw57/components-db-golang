schema "public" {
}

table "cabinets" {
  schema = schema.public

  column "id" {
    null = false
    type = bigserial
  }
  column "uuid" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "name" {
    null = false
    type = text
  }
  column "created_at" {
    null = false
    type = timestamptz
    default = sql("now()")
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_cabinets_uuid" {
    columns = [column.uuid]
  }
}
