package main

/**
Посетитель — это поведенческий паттерн проектирования,
который позволяет добавлять в программу новые операции, не изменяя классы объектов,
над которыми эти операции могут выполняться.
*/

// Visitor provides a visitor interface.
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Place интерфейс, которое посетитель должен посетить
/**
Метод принятия посетителя должен быть реализован в каждом
элементе, а не только в базовом классе. Это поможет программе
определить, какой метод посетителя нужно вызвать, если вы не
знаете тип элемента.
*/
type Place interface {
	Accept(v Visitor) string
}

// People реализует интерфейс Visitor.
type People struct {
}

// VisitSushiBar метод реализуется visit для SushiBar.
func (v *People) VisitSushiBar(place *SushiBar) string {
	return place.BuySushi()
}

// VisitPizzeria метод реализуется visit для Pizzeria.
func (v *People) VisitPizzeria(place *Pizzeria) string {
	return place.BuyPizza()
}

// VisitBurgerBar метод реализуется visit для BurgerBar.
func (v *People) VisitBurgerBar(place *BurgerBar) string {
	return place.BuyBurger()
}

// City implements a collection of places to visit.
type City struct {
	places []Place
}

// Add добавляет тип куда должен зайти Visitor
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept implements a visit to all places in the city.
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// SushiBar реализует Place
type SushiBar struct {
}

// Accept implementation.
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi implementation.
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Pizzeria реализует Place
type Pizzeria struct {
}

// Accept implementation.
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// BuyPizza implementation.
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// BurgerBar реализует Place
type BurgerBar struct {
}

// Accept implementation.
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// BuyBurger implementation.
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

/**
Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
не изменяя типы этих объектов, над которыми эти операции могут выполняться.

Нужен:
	Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов
	Когда новое поведение нужно только для некоторых классов из существующих

++ Плюсы
1. Упрощает добавление операций, работающих со сложными структурами объектов.
2. Посетитель может накапливать состояние при обходе структуры элементов

-- Минусы
1. Паттерн не оправдан, если иерархия элементов часто меняется
*/

func main() {
	people := &People{}
	places := []Place{&SushiBar{}, &Pizzeria{}, &BurgerBar{}}
	c := City{places: places}

	c.Accept(people)
}
