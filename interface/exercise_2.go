package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct{}

func (en EmailNotifier) Send(message string) error {
	fmt.Printf("Sending Email: %s\n", message)
	return nil
}

type SMSNotifier struct{}

func (sn SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS: %s\n", message)
	return nil
}

type NotificationLog struct {
	Recipient string
	Message   string
	Success   bool
}

func DispatchNotification(n Notifier, recipient string, message string) NotificationLog {
	err := n.Send(message)

	// Accept interface (n), Return concrete struct (NotificationLog)
	return NotificationLog{
		Recipient: recipient,
		Message:   message,
		Success:   err == nil, // True if no error occurred
	}
}

func main() {
	// Call using EmailNotifier
	emailLog := DispatchNotification(EmailNotifier{}, "faith@example.com", "Hi, Zander. This is Faith")
	fmt.Printf("Log: %+v\n\n", emailLog)

	// Call using SMSNotifier
	smsLog := DispatchNotification(SMSNotifier{}, "12345", "Hi, Zander. This is Faith via SMS")
	fmt.Printf("Log: %+v\n", smsLog)
}
