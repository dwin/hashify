package main

import (
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	r := gofight.New()
	r.GET("/status").SetDebug(true).Run(Router(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
	})
}
