// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package service

import (
	"context"
	"dynamodb-local-test/pkg/model"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

var ctx = context.TODO()

type PostService interface {
	Add(post model.Post) (string, error)
	Get(id string) (*model.Post, error)
}

type DdbPostService struct {
	ddbSvc    *dynamodb.Client
	tableName string
}

func (ps DdbPostService) Get(id string) (*model.Post, error) {

	var post model.Post

	input := &dynamodb.GetItemInput{
		TableName: aws.String("blog-post-table"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		}}

	res, err := ps.ddbSvc.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	if len(res.Item) == 0 {
		return nil, errors.Errorf("Record with ID %v is not found", id)
	}

	err = attributevalue.UnmarshalMap(res.Item, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (ps DdbPostService) Add(post model.Post) (string, error) {

	log.Printf("Post to add: %+v\n", post)

	av, err := attributevalue.MarshalMap(post)
	if err != nil {
		return "", errors.Wrapf(err, "error marshalling post info for DDB")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(ps.tableName),
	}

	_, err = ps.ddbSvc.PutItem(ctx, input)

	if err != nil {
		return "", errors.Wrapf(err, "error adding post DDB: %#v", post)
	}

	return "", nil
}

func NewDdbPostService(ddbSvc *dynamodb.Client) (PostService, error) {
	return &DdbPostService{
		ddbSvc:    ddbSvc,
		tableName: "blog-post-table",
	}, nil
}
