package conv

import (
	"bytes"
	"testing"
)

func TestBytesToString(t *testing.T) {
	data := []byte{'a', 'b', 'c'}

	expected := "abc"
	actual := BytesToString(data)

	if actual != expected {
		t.Errorf(`Expected - %s; actual - %s`, expected, actual)
	}
}

func TestStringToBytes(t *testing.T) {
	data := "abc"

	expected := []byte{'a', 'b', 'c'}
	actual := StringToBytes(data)

	if !bytes.Equal(expected, actual) {
		t.Errorf(`Expected - %v; actual - %v`, expected, actual)
	}
}
