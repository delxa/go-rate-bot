package main

import (
  "testing"
)

func TesContains(t *testing.T) {
  st := []string{"3000", "3001", "2000"}

  if Contains(st, "3000") != true {
  	t.Error("Array to contain 3000")
  }
  if Contains(st, "2322") == true {
  	t.Error("Expected array to not contain 2322")
  }
}