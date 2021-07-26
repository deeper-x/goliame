package goliame

import (
	"testing"
)

func TestLoadFile(t *testing.T) {
	m := New("demo", "demo", "host", "587", []string{"albertodeprezzo@gmail.com"})

	err := m.LoadFile("./assets/demo.txt")

	if err != nil {
		t.Error(err)
	}

	if string(m.Body) != "demo text" {
		t.Error("body content reading error")
	}
}
