package main

import (
	"testing"

	"gorm.io/gorm"
)

func TestConnection(t *testing.T) {
	t.Run("testing a wrong db uri", func(t *testing.T) {
		got := ConnectDB("user:root@tcp(localhost:3306)/moshdb")
		want := *gorm.DB

		if got != want {
			t.Fatal("got %v, want %v", got, want)
		}
	})
}
