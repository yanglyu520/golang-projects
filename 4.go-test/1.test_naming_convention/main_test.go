package main

import "testing"

func TestColor (t *testing.T) {
  arg := "blue"
  want := "#0000FF"
  got := Color(arg)
  if got != want {
    t.Errorf("color(%q)= %q, want %q", arg, got, want)
  }
}