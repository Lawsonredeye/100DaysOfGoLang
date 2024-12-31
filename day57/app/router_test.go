package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateTask(t *testing.T) {
	w := httptest.NewRecorder()
	params := bytes.NewBuffer([]byte(`{"title": "code at night"}`))
	req := httptest.NewRequest(http.MethodPost, "/tasks", params)
	tc, _ := gin.CreateTestContext(w)
	tc.Request = req
	CreateTask(tc)

	re := w.Result()

	if re.StatusCode != http.StatusCreated {
		t.Errorf("expected %v, got %v", http.StatusCreated, re.StatusCode)
	}
}

func TestGetTaskData(t *testing.T) {
	newWriter := httptest.NewRecorder()
	tc, _ := gin.CreateTestContext(newWriter)
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	tc.Request = req
	GetTask(tc)

	body := &Task{}
	_ = json.NewDecoder(newWriter.Body).Decode(&body)
	//_ = tc.BindJSON(&body)

	if body.Title != "code at night" {
		t.Errorf("expected %v, got %v", "code at night", body.Title)
	}

}
