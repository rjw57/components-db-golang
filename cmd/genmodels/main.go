package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

func main() {
	dsn := flag.String("dsn", "", "Postgres connection DSN")
	destDir := flag.String("destDir", "./gen", "Generated code directory")
	schema := flag.String("schema", "public", "Schema to generate models for")

	flag.Parse()

	if *dsn == "" {
		fmt.Fprintln(os.Stderr, "DSN must not be blank")
		os.Exit(1)
	}

	fmt.Printf("Using DSN: %v\n", *dsn)
	fmt.Printf("Examining schema: %v\n", *schema)
	fmt.Printf("Writing to: %v\n", *destDir)

	var fakeSkipFields = []string{"id", "created_at", "updated_at"}
	var skipTables = []string{"databasechangelog", "databasechangeloglock"}

	shouldSkipTable := func(table metadata.Table) bool {
		return stringSliceContains(skipTables, strings.ToLower(table.Name))
	}

	shouldFakeSkipField := func(columnMetaData metadata.Column) bool {
		return stringSliceContains(fakeSkipFields, strings.ToLower(columnMetaData.Name))
	}

	err := postgres.GenerateDSN(*dsn, *schema, *destDir,
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							if shouldSkipTable(table) {
								fmt.Printf("Skipping table: %v\n", table.Name)
								return template.TableModel{Skip: true}
							}
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									if shouldFakeSkipField(columnMetaData) {
										defaultTableModelField = defaultTableModelField.UseTags(`fake:"skip"`)
									}
									return defaultTableModelField
								})
						}),
					).
					UseSQLBuilder(template.DefaultSQLBuilder().
						UseTable(func(table metadata.Table) template.TableSQLBuilder {
							if shouldSkipTable(table) {
								fmt.Printf("Skipping table: %v\n", table.Name)
								return template.TableSQLBuilder{Skip: true}
							}
							return template.DefaultTableSQLBuilder(table)
						}),
					)
			}),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func stringSliceContains(ss []string, s string) bool {
	for _, a := range ss {
		if a == s {
			return true
		}
	}
	return false
}
