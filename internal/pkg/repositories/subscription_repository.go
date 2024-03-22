package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jrolstad/blog-monitor/internal/configuration"
	"github.com/jrolstad/blog-monitor/internal/pkg/clients"
	"github.com/jrolstad/blog-monitor/internal/pkg/models"
	"strconv"
)

type SubscriptionRepository interface {
	GetSubscriptions() ([]*models.Subscription, error)
}

func NewSubscriptionRepository(config *configuration.AppConfig) SubscriptionRepository {
	awsSession := clients.GetAwsSession(config.AwsRegion)
	return &DynamoDbSubscriptionRepository{
		client:    dynamodb.New(awsSession),
		tableName: config.SubscriptionTableName,
	}
}

type DynamoDbSubscriptionRepository struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func (r *DynamoDbSubscriptionRepository) GetSubscriptions() ([]*models.Subscription, error) {
	result := make([]*models.Subscription, 0)
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}
	queryResult, err := r.client.Scan(scanInput)
	if err != nil {
		return result, err
	}

	for _, item := range queryResult.Items {
		mappedItem := r.mapItemToSubscription(item)
		result = append(result, mappedItem)
	}

	return result, nil
}

func (r *DynamoDbSubscriptionRepository) mapItemToSubscription(item map[string]*dynamodb.AttributeValue) *models.Subscription {
	return &models.Subscription{
		Id:                  r.getStringValue(item["id"]),
		Name:                r.getStringValue(item["name"]),
		BlogUrl:             r.getStringValue(item["url"]),
		NotificationMethod:  r.getStringValue(item["notificationMethod"]),
		NotificationTargets: r.getStringSliceValue(item["notificationTargets"]),
		MaximumLookback:     r.getIntValue(item["maximumLookback"]),
	}
}

func (r *DynamoDbSubscriptionRepository) getStringValue(item *dynamodb.AttributeValue) string {
	if item.S == nil {
		return ""
	}
	return aws.StringValue(item.S)
}

func (r *DynamoDbSubscriptionRepository) getStringSliceValue(item *dynamodb.AttributeValue) []string {
	if item.SS == nil {
		return make([]string, 0)
	}
	return aws.StringValueSlice(item.SS)
}

func (r *DynamoDbSubscriptionRepository) getIntValue(item *dynamodb.AttributeValue) int {
	if item.N == nil {
		return 0
	}
	value := aws.StringValue(item.N)
	result, _ := strconv.Atoi(value)

	return result
}
