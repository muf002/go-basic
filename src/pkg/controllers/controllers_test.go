package controllers

import (
	"fmt"
	"testing"
)

func TestGetAllEmployees(t *testing.T) {
	want := "hello world"
	fmt.Println(want)
	if want != "hello world" {
		t.Fatalf(`want = "flan", want match for %#q, nil`, want)
	}
}
