env "local" {
  src = "file://schema.hcl"
  url = "postgresql://postgres:postgres-pass@db:5432/components?sslmode=disable"
}
