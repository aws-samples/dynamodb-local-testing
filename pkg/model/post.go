// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package model

type Post struct {
	Id                  string `dynamodbav:"id" json:"id"`
	Title               string `dynamodbav:"title" json:"title"`
	Content             string `dynamodbav:"content" json:"content"`
	Status              string `dynamodbav:"status" json:"status"`
	CreateTimestamp     string `dynamodbav:"createTimestamp" json:"create_timestamp"`
	LastUpdateTimestamp string `dynamodbav:"lastUpdateTimestamp" json:"update_timestamp"`
}
