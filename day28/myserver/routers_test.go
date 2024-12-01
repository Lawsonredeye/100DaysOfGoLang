package myserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("test login credentials", func(t *testing.T) {
		req := struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}{
			Name:     "lawson",
			Password: "testingroute",
		}

		params, _ := json.Marshal(req)
		got, err := http.Post("localhost:8080/login", "application/json", bytes.NewReader(params))
		if err != nil {
			t.Errorf("could not send request..")
		}
		defer got.Body.Close()
		// want := 200
		fmt.Println(got)

		// if got.StatusCode != want {
		// 	t.Errorf("wrong credentials")
		// }
	})
}
