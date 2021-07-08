// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package api

import (
	"dynamodb-local-test/pkg/model"
	"dynamodb-local-test/pkg/service"
	"dynamodb-local-test/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostServiceApi struct {
	PostService service.PostService
}

func (h *PostServiceApi) PostServiceGetApi(c *gin.Context) {
	id := c.Param("id")
	post, err := h.PostService.Get(id)
	if handleError400(c, err) {
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostServiceApi) PostServiceAddApi(c *gin.Context) {

	var post model.Post

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if handleError500(c, err) {
		return
	}

	err = json.Unmarshal(jsonData, &post)
	if handleError500(c, err) {
		return
	}

	if post.CreateTimestamp == "" {
		post.CreateTimestamp = utils.GetLocalTimestampNow()
	}

	if post.LastUpdateTimestamp == "" {
		post.LastUpdateTimestamp = utils.GetLocalTimestampNow()
	}

	_, err = h.PostService.Add(post)
	if handleError500(c, err) {
		return
	}

	c.JSON(http.StatusOK, post)
}

func handleError500(c *gin.Context, err error) bool {
	return handleError(c, err, http.StatusInternalServerError)
}

func handleError400(c *gin.Context, err error) bool {
	return handleError(c, err, http.StatusBadRequest)
}

func handleError(c *gin.Context, err error, status int) bool {
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return true
	}
	return false
}
