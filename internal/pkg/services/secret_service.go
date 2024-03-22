package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/jrolstad/blog-monitor/internal/pkg/clients"
	"github.com/jrolstad/blog-monitor/internal/pkg/configuration"
)

type SecretService interface {
	GetGoogleApiKey() (string, error)
}

func NewSecretService(config *configuration.AppConfig) SecretService {
	awsSession := clients.GetAwsSession(config.AwsRegion)
	return &SecretsManagerSecretService{
		client:                 secretsmanager.New(awsSession),
		googleApiKeySecretName: config.GoogleApiKeySecretName,
	}
}

type SecretsManagerSecretService struct {
	client                 *secretsmanager.SecretsManager
	googleApiKeySecretName string
}

func (s *SecretsManagerSecretService) GetGoogleApiKey() (string, error) {
	return s.getValue(s.googleApiKeySecretName)
}

func (s *SecretsManagerSecretService) getValue(name string) (string, error) {
	return s.getSecret(name)
}

func (s *SecretsManagerSecretService) getSecret(name string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}

	secretValue, err := s.client.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	return *secretValue.SecretString, nil
}
