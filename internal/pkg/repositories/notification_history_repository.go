package repositories

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jrolstad/blog-monitor/internal/pkg/clients"
	"github.com/jrolstad/blog-monitor/internal/pkg/configuration"
	"time"
)

type NotificationHistoryRepository interface {
	Exists(subscriptionId string, postId string) (bool, error)
	TrackNotification(subscriptionId string, postId string, notifiedAt time.Time) error
}

func NewNotificationHistoryRepository(config *configuration.AppConfig) NotificationHistoryRepository {
	awsSession := clients.GetAwsSession(config.AwsRegion)
	return &DynamoDbNotificationHistoryRepository{
		client:    dynamodb.New(awsSession),
		tableName: config.NotificationHistoryTableName,
	}
}

type DynamoDbNotificationHistoryRepository struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func (s *DynamoDbNotificationHistoryRepository) Exists(subscriptionId string, postId string) (bool, error) {
	itemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(createIdentifier(subscriptionId, postId)),
			},
		},
		TableName: aws.String(s.tableName),
	}
	queryResult, err := s.client.GetItem(itemInput)
	if err != nil {
		return false, err
	}

	exists := queryResult.Item != nil
	return exists, nil
}

func (s *DynamoDbNotificationHistoryRepository) TrackNotification(subscriptionId string, postId string, notifiedAt time.Time) error {

	data := make(map[string]*dynamodb.AttributeValue)
	data["id"] = &dynamodb.AttributeValue{S: aws.String(createIdentifier(subscriptionId, postId))}
	data["subscription"] = &dynamodb.AttributeValue{S: aws.String(subscriptionId)}
	data["post"] = &dynamodb.AttributeValue{S: aws.String(postId)}
	data["notifiedAt"] = &dynamodb.AttributeValue{S: aws.String(notifiedAt.String())}

	request := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(s.tableName),
	}
	_, err := s.client.PutItem(request)

	return err
}

func createIdentifier(subscriptionId string, postId string) string {
	return fmt.Sprintf("%s|%s", subscriptionId, postId)
}
