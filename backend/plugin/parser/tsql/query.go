package tsql

import (
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"

	"github.com/bytebase/bytebase/backend/plugin/parser/base"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

func init() {
	base.RegisterExtractResourceListFunc(storepb.Engine_MSSQL, ExtractResourceList)
	base.RegisterExtractDatabaseListFunc(storepb.Engine_MSSQL, ExtractDatabaseList)
}

// ExtractDatabaseList extracts the database names.
func ExtractDatabaseList(statement string, normalizedDatabaseName string) ([]string, error) {
	schemaPlaceholder := "dbo"
	schemaResource, err := ExtractResourceList(normalizedDatabaseName, schemaPlaceholder, statement)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, resource := range schemaResource {
		result = append(result, resource.Database)
	}
	return result, nil
}

// ExtractResourceList extracts the list of resources from the SELECT statement, and normalizes the object names with the NON-EMPTY currentNormalizedDatabase and currentNormalizedSchema.
func ExtractResourceList(currentNormalizedDatabase string, currentNormalizedSchema string, selectStatement string) ([]base.SchemaResource, error) {
	tree, err := ParseTSQL(selectStatement)
	if err != nil {
		return nil, err
	}

	l := &tsqlReasourceExtractListener{
		currentDatabase: currentNormalizedDatabase,
		currentSchema:   currentNormalizedSchema,
		resourceMap:     make(map[string]base.SchemaResource),
	}

	var result []base.SchemaResource
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	for _, resource := range l.resourceMap {
		result = append(result, resource)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].String() < result[j].String()
	})

	return result, nil
}

type tsqlReasourceExtractListener struct {
	*parser.BaseTSqlParserListener

	currentDatabase string
	currentSchema   string
	resourceMap     map[string]base.SchemaResource
}

// EnterTable_source_item is called when the parser enters the table_source_item production.
func (l *tsqlReasourceExtractListener) EnterTable_source_item(ctx *parser.Table_source_itemContext) {
	if fullTableName := ctx.Full_table_name(); fullTableName != nil {
		var parts []string
		var linkedServer string
		if server := fullTableName.GetLinkedServer(); server != nil {
			linkedServer = NormalizeTSQLIdentifier(server)
		}
		parts = append(parts, linkedServer)

		database := l.currentDatabase
		if d := fullTableName.GetDatabase(); d != nil {
			normalizedD := NormalizeTSQLIdentifier(d)
			if normalizedD != "" {
				database = normalizedD
			}
		}
		parts = append(parts, database)

		schema := l.currentSchema
		if s := fullTableName.GetSchema(); s != nil {
			normalizedS := NormalizeTSQLIdentifier(s)
			if normalizedS != "" {
				schema = normalizedS
			}
		}
		parts = append(parts, schema)

		var table string
		if t := fullTableName.GetTable(); t != nil {
			normalizedT := NormalizeTSQLIdentifier(t)
			if normalizedT != "" {
				table = normalizedT
			}
		}
		parts = append(parts, table)
		normalizedObjectName := strings.Join(parts, ".")
		l.resourceMap[normalizedObjectName] = base.SchemaResource{
			LinkedServer: linkedServer,
			Database:     database,
			Schema:       schema,
			Table:        table,
		}
	}

	if rowsetFunction := ctx.Rowset_function(); rowsetFunction != nil {
		return
	}

	// https://simonlearningsqlserver.wordpress.com/tag/changetable/
	// It seems that the CHANGETABLE is only return some statistics, so we ignore it.
	if changeTable := ctx.Change_table(); changeTable != nil {
		return
	}

	// other...
}
