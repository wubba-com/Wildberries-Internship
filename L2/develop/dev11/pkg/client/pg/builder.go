package pg

import (
	"fmt"
	build "github.com/wubba-com/L2/develop/dev11/pkg/client/builder"
	"strings"
)

func NewPSGQueryBuilder() build.SQLQueryBuilder {
	return &PSGSQLQueryBuilder{}
}

// PSGSQLQueryBuilder Конкретный Строитель соответствует определённому диалекту SQL и может
// реализовать шаги построения немного иначе, чем остальные.
type PSGSQLQueryBuilder struct {
	Query string
}

func (b *PSGSQLQueryBuilder) Get() string {
	query := b.Query
	b.Query = ""
	return query
}

func (b *PSGSQLQueryBuilder) Select(table string, fields []string) build.SQLQueryBuilder {
	b.Query = fmt.Sprintf("SELECT %s FROM %s ", strings.Join(fields, ", "), table)
	return b
}

func (b *PSGSQLQueryBuilder) Where(field string, operator string, value string) build.SQLQueryBuilder {
	b.Query += fmt.Sprintf("WHERE %s %s %s ", field, operator, value)
	return b
}

func (b *PSGSQLQueryBuilder) WhereAnd(field, operator, value, field2, operator2, value2 string) build.SQLQueryBuilder {
	b.Query += fmt.Sprintf("WHERE %s %s %s AND %s %s %s ", field, operator, value, field2, operator2, value2)
	return b
}

func (b *PSGSQLQueryBuilder) Returning(field string) build.SQLQueryBuilder {
	b.Query += fmt.Sprintf("RETURNING %s ", field)
	return b
}

func (b *PSGSQLQueryBuilder) OrderBy(field, how string) build.SQLQueryBuilder {
	b.Query += fmt.Sprintf("ORDER BY %s %s ", field, strings.ToUpper(how))
	return b
}

func (b *PSGSQLQueryBuilder) Limit(limit int) build.SQLQueryBuilder {
	b.Query += fmt.Sprintf("LIMIT %d ", limit)
	return b
}

func (b *PSGSQLQueryBuilder) Insert(table string, fields []string) build.SQLQueryBuilder {
	var values string
	var i = 1
	for len(fields) >= i {
		if len(fields) == i {
			values += fmt.Sprintf("$%d", i)
			break
		}
		values += fmt.Sprintf("$%d, ", i)
		i++
	}
	b.Query += fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ", table, strings.Join(fields, ", "), values)

	return b
}

func (b *PSGSQLQueryBuilder) Delete(table, field string) build.SQLQueryBuilder {
	b.Query += fmt.Sprintf("DELETE FROM %s WHERE %s = $1 ", table, field)

	return b
}
