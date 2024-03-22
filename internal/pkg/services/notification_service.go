package services

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/jrolstad/blog-monitor/internal/configuration"
	"github.com/jrolstad/blog-monitor/internal/pkg/clients"
)

type NotificationService interface {
	Notify(method string, targets []string, title string, content string) error
}

type AmazonSesNotificationService struct {
	sesService *ses.SES
	sender     string
}

func NewNotificationService(config *configuration.AppConfig) NotificationService {
	awsSession := clients.GetAwsSession(config.AwsRegion)
	return &AmazonSesNotificationService{
		sesService: ses.New(awsSession),
		sender:     config.EmailSender,
	}
}

func (s *AmazonSesNotificationService) Notify(method string, targets []string, title string, content string) error {
	switch method {
	case "email":
		return s.sendEmail(targets, title, content)
	default:
		return errors.New("unsupported notification type")
	}
}

func (s *AmazonSesNotificationService) sendEmail(sendTo []string, title string, content string) error {
	input := s.mapToEmailMessage(sendTo, title, content)
	_, err := s.sesService.SendEmail(input)

	return err
}

func (s *AmazonSesNotificationService) mapToEmailMessage(sendTo []string, title string, content string) *ses.SendEmailInput {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: aws.StringSlice(sendTo),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(content),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(title),
			},
		},
		Source: aws.String(s.sender),
	}
	return input
}
