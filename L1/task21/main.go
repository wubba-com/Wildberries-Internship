package main

import "fmt"

/**
Реализовать паттерн «адаптер» на любом примере
*/

// Notification - интерфейс предоставляет интерфейс, которому следует приложение
type Notification interface {
	Send(title, msg string)
}

// EmailNotification - пример существующего класса, который следует за интерфейсом в приложении.
type EmailNotification struct {
	email string
}

// Send - EmailNotification реализует интерфейс, отправляя уведомление на почту
func (e *EmailNotification) Send(title, msg string) {
	fmt.Printf("Отправка письма на почту:%s с заголовком:%s, msg:%s \n", e.email, title, msg)
}

// SlackApi Сторонняя библиотека для работы с API Slack. Теперь мы хотим, что б уведомление приходили не только на почту
type SlackApi struct {
	login  string
	ApiKey string
}

// Login - Отправляет запрос аутентификации в Slack
func (s *SlackApi) Login() {
	fmt.Printf("Вы авторизовались в аккаунте '%s'\n", s.login)
}

// SendMessage - Отправляет запрос на публикацию сообщения в Slack
func (s *SlackApi) SendMessage(chatID, msg string) {
	fmt.Printf("Пользователь разместил сообщение в чате '%s' с сообщением: '%s'.\n", chatID, msg)
}

// SlackNotification Адаптер – класс, который связывает интерфейс использующий в приложении со сторонней библиотекой.
type SlackNotification struct {
	Slack  *SlackApi // Адаптер содержит в себе объект который нужно адаптировать под свои нужды
	ChatID string
}

func (s *SlackNotification) Send(title, msg string) {
	slackMSG := fmt.Sprintf("%s#%s", title, msg)
	s.Slack.Login()
	s.Slack.SendMessage(s.ChatID, slackMSG)
}

func App(notification Notification) {
	title := "pattern Adapter"
	msg := "Адаптер - это паттерн обертка, который позволяет адаптировать например, сторонюю библиотеку к формату использующем в своем приложении"
	notification.Send(title, msg)
}

func main() {

	// email
	emailNoti := &EmailNotification{email: "dev@gmail.com"}
	App(emailNoti)

	// slack
	slackApi := &SlackApi{ApiKey: "test123", login: "dev"}
	slackNoti := &SlackNotification{Slack: slackApi, ChatID: "someChatID"}

	App(slackNoti)
}
