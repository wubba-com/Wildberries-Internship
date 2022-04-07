package main

import (
	"errors"
	"fmt"
)

// Payment - Интерфейс Стратегии описывает, как клиент может использовать различные конкретные Стратегии
type Payment interface {
	Pay() error
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

type payPalPayment struct {
	cardNum string
	cvv     string
}

func (p payPalPayment) Pay() error {
	// API PayPal
	fmt.Println("paypal payment")
	return nil
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

type qiwiPayment struct {
	cardNum string
	cvv     string
}

func (q qiwiPayment) Pay() error {
	// API QIWI
	fmt.Println("QIWI payment")
	return nil
}

func NewSberPayPayment() Payment {
	return &SberPayPayment{}
}

type SberPayPayment struct {
	cardNum string
	cvv     string
}

func (c SberPayPayment) Pay() error {
	// API SberPay
	fmt.Println("SberPay payment")
	return nil
}

/**
Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов
и помещает каждый из них в собственный структуру, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
*/

/**
Нужен:
Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
Когда у вас есть множество похожих классов, отличающихся только некоторым поведением
Когда ветка такого оператора представляет собой вариацию алгоритма
*/

/**
++ Плюсы
1. Изолирует код и данные алгоритмов от остальных классов.
2. Реализует принцип открытости/закрытости.
3. Делегируем работу

-- Минусы
1. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую
*/

const (
	_ = iota
	SberPay
	PayPal
	QIWIPay
)

// PaymentFactory - фабрика создает объект для оплаты заказа
type PaymentFactory struct {
}

func (pf *PaymentFactory) PaymentMethod(choicePay int) Payment {
	switch choicePay {
	case SberPay:
		return NewSberPayPayment()
	case PayPal:
		return NewPayPalPayment()
	case QIWIPay:
		return NewQIWIPayment()
	}

	return nil
}

// Order структура заказа
type Order struct {
	OrderUID   string
	PaymentUID int
}

// FromForm данные из формы
func FromForm() (string, int) {
	return "uid-123", 3
}

// ProcessOrder - обработка заказа
func ProcessOrder(order *Order) error {
	pf := &PaymentFactory{}
	payment := pf.PaymentMethod(order.PaymentUID)
	if payment == nil {
		return errors.New("err payment")
	}

	err := payment.Pay()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	orderUID, paymentID := FromForm()
	order := &Order{orderUID, paymentID}

	err := ProcessOrder(order)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
