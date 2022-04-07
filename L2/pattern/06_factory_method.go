package main

const (
	postgresDB = "postgres"
	mysqlDB    = "mysql"
	mongoDB    = "mongo"
)

// FactoryClient Фабрика создающая клиент к бд в зависимости от переданного параметра из конфига или cli
func FactoryClient(client string) Client {
	switch client {
	case postgresDB:
		return NewPostgres()
	case mysqlDB:
		return NewMysql()
	case mongoDB:
		return NewMongo()
	}

	return nil
}

// Client Интерфейс продукта
type Client interface {
	Accept() error
}

// Подключение к Postgres.

func NewPostgres() Client {
	return &PostgresClient{}
}

type PostgresClient struct {
}

func (p *PostgresClient) Accept() error {
	return nil
}

// Подключение к Mysql.

// NewMysql новый объект Mysql
func NewMysql() Client {
	return &MysqlClient{}
}

type MysqlClient struct {
}

func (p *MysqlClient) Accept() error {
	return nil
}

// Подключение к Mongo.

func NewMongo() Client {
	return &MongoClient{}
}

type MongoClient struct {
}

func (p *MongoClient) Accept() error {
	return nil
}

/**
Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов.
Определяет интерфейс для создания объектов, но ост
Нужен:
	Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код

++ Плюсы
1.  Избавляет класс от привязки к конкретным классам продуктов
2.  Выделяет код производства продуктов в одно место, упрощая поддержку кода
3. Упрощает добавление новых структур в программу
4. Реализует принцип открытости/закрытости
*/
func main() {
	// параметр из конфига
	var dbConfig = "postgres"

	client := FactoryClient(dbConfig)
	err := client.Accept()
	if err != nil {
		return
	}
}
