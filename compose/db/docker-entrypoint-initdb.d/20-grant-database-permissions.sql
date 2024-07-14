GRANT TEMPORARY, CONNECT ON DATABASE backend TO "backend-user";
GRANT TEMPORARY, CONNECT ON DATABASE test TO "test-user";
GRANT "pg_read_all_data", "pg_write_all_data" TO "backend-user", "test-user";
