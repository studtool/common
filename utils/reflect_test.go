package utils

import (
	"testing"
)

type testType struct{}

func TestStructName(t *testing.T) {
	v := &testType{}

	actual := StructName(v)
	expected := "utils.testType"

	if actual != expected {
		t.Errorf(`expected %s buf actual is %s`, actual, expected)
	}
}
