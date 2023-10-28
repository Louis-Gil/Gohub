package db

import (
	"context"
	"errors"
	"server/domain/models"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const TableName = "Readme_Dev_Users"

type DynamoDBUserRepository struct {
	db *dynamodb.Client
}

func NewDynamoDBUserRepository(db *dynamodb.Client) *DynamoDBUserRepository {
	return &DynamoDBUserRepository{db: db}
}

func (r *DynamoDBUserRepository) Create(user *models.User) error {
	_, err := r.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item: map[string]types.AttributeValue{
			"id":         &types.AttributeValueMemberN{Value: strconv.Itoa(user.Id)},
			"first_name": &types.AttributeValueMemberS{Value: user.FirstName},
			"last_name":  &types.AttributeValueMemberS{Value: user.LastName},
			"email":      &types.AttributeValueMemberS{Value: user.Email},
		},
	})

	return err
}

func (r *DynamoDBUserRepository) GetByID(id int) (*models.User, error) {
	res, err := r.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
		},
	})

	if err != nil {
		return nil, err
	}

	if res.Item == nil {
		return nil, errors.New("user not found")
	}

	return &models.User{
		Id:        id,
		FirstName: res.Item["first_name"].(*types.AttributeValueMemberS).Value,
		LastName:  res.Item["last_name"].(*types.AttributeValueMemberS).Value,
		Email:     res.Item["email"].(*types.AttributeValueMemberS).Value,
	}, nil
}
