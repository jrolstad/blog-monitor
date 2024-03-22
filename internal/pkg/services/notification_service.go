package services

import (
	"errors"
)

type NotificationService interface {
	Notify(method string, targets []string, title string, content string) error
}

type AmazonSesNotificationService struct {
}

func NewNotificationService() NotificationService {
	return &AmazonSesNotificationService{}
}

func (s *AmazonSesNotificationService) Notify(method string, targets []string, title string, content string) error {
	switch method {
	case "email":
		s.sendEmail(targets, title, content)
	default:
		return errors.New("unsupported notification type")
	}

	return nil
}

func (*AmazonSesNotificationService) sendEmail(sendTo []string, title string, content string) {

}
