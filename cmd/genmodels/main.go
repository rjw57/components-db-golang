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

	var skipTables = []string{"databasechangelog", "databasechangeloglock"}
	var tableModelTypeNames = map[string]string{
		"cabinets": "Cabinet",
	}

	shouldSkipTable := func(table metadata.Table) bool {
		return stringSliceContains(skipTables, strings.ToLower(table.Name))
	}

	err := postgres.GenerateDSN(*dsn, *schema, *destDir,
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							if shouldSkipTable(table) {
								fmt.Printf("Table %s: skipping\n", table.Name)
								return template.TableModel{Skip: true}
							}
							tableModel := template.DefaultTableModel(table)
							if typeName, ok := tableModelTypeNames[table.Name]; ok {
								fmt.Printf("Table %s: using model type name %s\n", table.Name, typeName)
								tableModel = tableModel.UseTypeName(typeName)
							}
							return tableModel
						}),
					).
					UseSQLBuilder(template.DefaultSQLBuilder().
						UseTable(func(table metadata.Table) template.TableSQLBuilder {
							if shouldSkipTable(table) {
								fmt.Printf("Skipping table: %v\n", table.Name)
								return template.TableSQLBuilder{Skip: true}
							}
							tableSQLBuilder := template.DefaultTableSQLBuilder(table)
							if typeName, ok := tableModelTypeNames[table.Name]; ok {
								fmt.Printf("Table %s: using default alias %s\n", table.Name, typeName)
								tableSQLBuilder = tableSQLBuilder.UseDefaultAlias(typeName)
							}
							return tableSQLBuilder
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
