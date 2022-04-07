package main

import "fmt"

// Интерфейс Команды объявляет основной метод выполнения
type command interface {
	execute()
}

// Интерфейс получателя
type device interface {
	save()
	copy()
}

// Конкретная команда
type saveCommand struct {
	device device
}

func (c *saveCommand) execute() {
	c.device.save()
}

// copyCommand Конкретная команда
type copyCommand struct {
	device device
}

func (c *copyCommand) execute() {
	c.device.copy()
}

// Отправитель кнопка на пульте
type button struct {
	command command
}

// Действие в отправителе, нужно выполнить какую то работу
func (b *button) press() {
	b.command.execute()
}

func (b *button) setCommand(command command)  {
	b.command = command
}

// IDE Есть более сложные операции и команда может делегировать работу другим объектам - Получателям
type IDE struct {
	isRunning bool
}

func (i *IDE) save() {
	i.isRunning = true
	fmt.Println("save file")
}

func (i *IDE) copy() {
	i.isRunning = false
	fmt.Println("copy text")
}

/**
Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций

Нужен:
	Когда нужно передавать действие как объект
	Когда нужно ставить операции в очередь, выполнять их по расписанию или передавать по сети
	Когда вам нужна операция отмены

++ Плюсы
1. Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
2. Позволяет реализовать простую отмену и повтор операций
3. Позволяет реализовать отложенный запуск операций
4. Реализует принцип открытости/закрытости

-- Минусы
1. Усложняет код программы из-за введения множества дополнительных классов
 */

func main()  {
	ide := &IDE{}
	saveC := &saveCommand{device: ide}
	copyC := &copyCommand{device: ide}

	btn := &button{command: saveC}
	btn.press()

	btn.setCommand(copyC)
	btn.press()
}