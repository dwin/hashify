package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var testValue = "helloWorld"

func TestHashSHA1Get(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/sha1?value="+testValue, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HashSHA1(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "5395ebfd174b0a5617e6f409dfbb3e064e3fdf0a") // should equal SHA1 of "helloWorld"
	}
}

func TestHashSHA1Post(t *testing.T) {
	// Setup
	e := echo.New()
	body := new(bytes.Buffer)
	_, err := body.Write([]byte(testValue))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	req := httptest.NewRequest(echo.POST, "/sha1", body)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HashSHA1(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "5395ebfd174b0a5617e6f409dfbb3e064e3fdf0a") // should equal SHA1 of "helloWorld"
	}
}
