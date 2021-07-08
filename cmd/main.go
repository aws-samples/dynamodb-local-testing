// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package main

import (
	"context"
	"dynamodb-local-test/pkg/api"
	"dynamodb-local-test/pkg/model"
	"dynamodb-local-test/pkg/service"
	"dynamodb-local-test/pkg/utils"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

var awsCfg aws.Config
var ctx = context.TODO()

func main() {
	fmt.Println("Starting Server")

	post := model.Post{}
	post.Id = "1"
	post.Title = "my post"
	post.Content = "post content"
	post.Status = "posted"
	post.CreateTimestamp = utils.GetLocalTimestampNow()
	post.LastUpdateTimestamp = utils.GetLocalTimestampNow()

	awsProfile, ok := os.LookupEnv("AWS_PROFILE")
	log.Printf("AWS_PROFILE: %s", awsProfile)
	var err error
	var ps service.PostService

	if ok {
		log.Printf("Use AWS profile")
		awsCfg, err = config.LoadDefaultConfig(ctx,
			config.WithSharedConfigProfile(awsProfile),
		)
		if err != nil {
			log.Fatalf("Error loading profile %v", err)
		}

	} else {
		log.Printf("Use default role")
		awsCfg, err = config.LoadDefaultConfig(ctx)
	}

	ddbSvc := dynamodb.NewFromConfig(awsCfg)
	log.Printf("DDB service created")

	ps, _ = service.NewDdbPostService(ddbSvc)

	postServiceHandlers := api.PostServiceApi{PostService: ps}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	router.GET("/post/:id", postServiceHandlers.PostServiceGetApi)

	router.POST("/post", postServiceHandlers.PostServiceAddApi)

	_ = router.Run(":80")

}
