package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUser(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	GetUser(c)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %v\n", w.Code)
	}
}
