// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package service

import (
	"context"
	"dynamodb-local-test/pkg/model"
	"dynamodb-local-test/pkg/utils"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"

	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func newPostService() (PostService, error) {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == "us-east-1" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:8000",
				SigningRegion: "us-east-1",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ddbSvc := dynamodb.NewFromConfig(awsCfg)
	log.Printf("DDB service created")

	err = createTable(ddbSvc, "blog-post-table")
	log.Printf("DDB table created")

	ps, _ := NewDdbPostService(ddbSvc)
	return ps, nil

}

func createTable(ddbSvc *dynamodb.Client, tableName string) error {
	_, err := ddbSvc.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		}},
	)

	return err
}

func Test_PostService(t *testing.T) {

	post := model.Post{}
	post.Id = "1"
	post.Title = "my post"
	post.Content = "post content"
	post.Status = "posted"
	post.CreateTimestamp = utils.GetLocalTimestampNow()
	post.LastUpdateTimestamp = utils.GetLocalTimestampNow()

	ps, _ := newPostService()

	_, err := ps.Add(post)
	if err != nil {
		log.Fatalf("Error storing post: %v", err)
	}
	log.Printf("post added")

	post2, err := ps.Get("1")

	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Post from DDB: %+v", post2)

	assert.EqualValues(t, post, *post2)
}
