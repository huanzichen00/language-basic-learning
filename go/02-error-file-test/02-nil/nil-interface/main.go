package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
	Address string
}

func (e *EmailNotifier) Send(message string) {
	if e == nil {
		fmt.Println("email notifier is nil")
		return
	}
	fmt.Println("send email to", e.Address+":", message)
}

func sendNotification(notifier Notifier) {
	if notifier == nil {
		fmt.Println("notifier is nil")
		return
	}

	notifier.Send("hello")
}

func main() {
	var notifier Notifier
	fmt.Println("case 1:")
	sendNotification(notifier)

	// != nil
	// 携带 type = EmailNotifier 信息
	var email *EmailNotifier
	fmt.Println("case 2:")
	sendNotification(email)

	email = &EmailNotifier{
		Address: "tom@example.com",
	}

	fmt.Println("case 3:")
	sendNotification(email)
}
