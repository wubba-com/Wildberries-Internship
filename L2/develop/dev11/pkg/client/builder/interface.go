package build

// SQLQueryBuilder Интерфейс Строителя объявляет набор методов для сборки SQL-запроса
type SQLQueryBuilder interface {
	Select(table string, fields []string) SQLQueryBuilder
	Where(field string, operator string, value string) SQLQueryBuilder
	WhereAnd(field, operator, value, field2, operator2, value2 string) SQLQueryBuilder
	Returning(field string) SQLQueryBuilder
	OrderBy(field, how string) SQLQueryBuilder
	Limit(limit int) SQLQueryBuilder
	Insert(table string, fields []string) SQLQueryBuilder
	Delete(table, field string) SQLQueryBuilder
	Get() string
}
