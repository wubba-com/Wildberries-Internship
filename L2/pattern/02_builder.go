package main

import (
	"fmt"
	"strings"
)

// SQLQueryBuilder Интерфейс Строителя объявляет набор методов для сборки SQL-запроса
type SQLQueryBuilder interface {
	Select(table string, fields []string) SQLQueryBuilder
	Where(field string, operator string, value string) SQLQueryBuilder
	Limit(limit int) SQLQueryBuilder
	Get() string
}

func NewMySqlBuilder() SQLQueryBuilder {
	return &MySQLQueryBuilder{}
}

// MySQLQueryBuilder Конкретный Строитель соответствует определённому диалекту SQL и может
// реализовать шаги построения немного иначе, чем остальные.
type MySQLQueryBuilder struct {
	Query string
}

func (b *MySQLQueryBuilder) Get() string {
	return b.Query
}

func (b *MySQLQueryBuilder) Select(table string, fields []string) SQLQueryBuilder {
	b.Query = fmt.Sprintf("SELECT %s FROM %s ", strings.Join(fields, ", "), table)
	return b
}

func (b *MySQLQueryBuilder) Where(field string, operator string, value string) SQLQueryBuilder {
	b.Query += fmt.Sprintf("WHERE %s %s %s ", field, operator, value)
	return b
}

func (b *MySQLQueryBuilder) Limit(limit int) SQLQueryBuilder {
	b.Query += fmt.Sprintf("LIMIT %d", limit)
	return b
}

/**
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

// Строитель можно использовать в репозиториях для построения строки запроса

Нужен:
1. Когда вам нужно собирать сложные составные объекты. Строитель конструирует объекты пошагово, а не за один проход
2. Когда ваш код должен создавать разные представления какого-то объекта

++ Плюсы
Позволяет создавать продукты пошагово.
Позволяет использовать один и тот же код для создания различных продуктов.
Изолирует сложный код сборки продукта

-- Минусы
1. Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.

Какую проблему решает паттерн:
Есть сложный объект, требующий кропотливой пошаговой инициализации множества полей и вложенных объектов.
Код инициализации таких объектов состоит из 10 параметров
*/

func main() {
	concreteBuilder := NewMySqlBuilder()
	query := concreteBuilder.Select("user", []string{"name", "email"}).Where("id", "=", "1").Get()
	fmt.Println(query)
}
