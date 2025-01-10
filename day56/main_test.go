package main

import (
	"encoding/json"
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

type User struct {
	Name     string `json:"name"`
	CodeName string `json:"codename"`
}

func TestGetUserData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	GetUser(c)
	// body, _ := io.ReadAll(w.Body)
	user := User{}
	_ = json.NewDecoder(w.Body).Decode(&user)

	if user.Name != "lawsonredeye" {
		t.Errorf("expected lawsonredeye, got %v\n", user.Name)
	}
}
