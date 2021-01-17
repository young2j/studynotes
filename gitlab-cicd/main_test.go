package main

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_queryUser(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"", "c君"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryUser(); got != tt.want {
				t.Errorf("queryUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
