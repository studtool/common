package rft

import (
	"testing"
)

type testType struct{}

func testFunc() {

}

func TestFuncName(t *testing.T) {
	actual := FuncName(testFunc)
	expected := "rft.testFunc"

	if actual != expected {
		t.Errorf(`expected %s buf actual is %s`, actual, expected)
	}
}

func TestStructName(t *testing.T) {
	v := &testType{}

	actual := StructName(v)
	expected := "rft.testType"

	if actual != expected {
		t.Errorf(`expected %s buf actual is %s`, actual, expected)
	}
}
