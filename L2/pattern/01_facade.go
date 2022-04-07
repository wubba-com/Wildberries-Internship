package main
import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func NewValidaterFacade() *validaterFacade {
	v := validator.New()
	return &validaterFacade{validater: v}
}

type validaterFacade struct {
	validater *validator.Validate
}

func (v *validaterFacade) ValidateStruct(val interface{}) error {
	err := v.validater.Struct(val)
	if err != nil {
		return err

	}
	return nil
}

func (v *validaterFacade) IsEmail(email string) error {
	err := v.validater.Var(email, "required,email")
	if err != nil {
		return err
	}
	return nil
}

// Фасад - это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.

/**
Какую проблему решает паттерн:
Когда в приложении приходится работать с большим количеством объектов библиотеки или с огромным количеством методов структур,
которые не нужны вначале или вообще или когда хочется ограничить какой-то функционал фреймворка или упростить его.
Когда нам приходиться самостоятельно инициализировать эти объекты и следить за правильным порядком зависимостей итд
*/

type User struct {
	UserUID string `validate:"required"`
	Email   string `validate:"required,email"`
	Name    string `validate:"required"`
	Age     uint8  `validate:"gte=0,lte=130"`
}

func main() {
	v := NewValidaterFacade()
	u := &User{UserUID: "1", Email: "test", Name: "Petya", Age: 25}
	err := v.IsEmail(u.Email)
	if err != nil {
		fmt.Println("v email", err.Error())
	}

	err = v.ValidateStruct(u)
	if err != nil {
		fmt.Println("v struct", err.Error())
		return
	}
}


