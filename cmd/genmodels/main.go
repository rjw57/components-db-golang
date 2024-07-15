package main

import (
	"flag"
	"fmt"
	"os"

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

	err := postgres.GenerateDSN(*dsn, *schema, *destDir,
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									n := columnMetaData.Name
									if n == "id" || n == "created_at" || n == "updated_at" {
										defaultTableModelField = defaultTableModelField.UseTags(
											`fake:"skip"`,
										)
									}
									return defaultTableModelField
								})
						}),
					)
			}),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
