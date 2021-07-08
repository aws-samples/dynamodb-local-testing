// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0

package utils

import "time"

func GetLocalTimestamp(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

func GetLocalTimestampNow() string {
	t := time.Now()
	return GetLocalTimestamp(t)
}
