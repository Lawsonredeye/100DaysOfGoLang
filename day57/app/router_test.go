package router

import (
	"bytes"
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
