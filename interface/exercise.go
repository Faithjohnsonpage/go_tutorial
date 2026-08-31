package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (en EmailNotifier) Send(message string) {
	fmt.Printf("Sending Email: %s\n", message)
}

type SMSNotifier struct{}

func (sn SMSNotifier) Send(message string) {
	fmt.Printf("Sending SMS: %s\n", message)
}

type NotificationService struct {
	N Notifier
}

func (ns NotificationService) Notify(msg string) {
	ns.N.Send(msg)
}

func main() {
	// Initialize with EmailNotifier
	n := NotificationService{
		N: EmailNotifier{},
	}
	n.Notify("Hello!")

	// Swap implementation to SMSNotifier
	n = NotificationService{
		N: SMSNotifier{},
	}
	n.Notify("Hello!")
}
