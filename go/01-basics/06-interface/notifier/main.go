package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Println("send email to", e.Address+":", message)
	return nil
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(message string) error {
	fmt.Println("send sms to", s.Phone+":", message)
	return nil
}

func notify(notifier Notifier, message string) {
	err := notifier.Send(message)
	if err != nil {
		fmt.Println("notify error:", err)
		return
	}

	fmt.Println("notify success")
}

func main() {
	email := EmailNotifier{
		Address: "tom@example.com",
	}

	sms := SMSNotifier{
		Phone: "13800000000",
	}

	notify(email, "Your order has been shipped")
	notify(sms, "Your code is 123456")
}
