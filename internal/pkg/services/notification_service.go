package services

import "fmt"

type NotificationService struct {
}

func NewNotificationService() NotificationService {
	return NotificationService{}
}

func (*NotificationService) Notify(method string, targets []string, title string, content string) error {
	fmt.Println(title)
	fmt.Println(content)
	return nil
}
