package rft

import (
	"testing"
)

func testFunc() {}

type testType struct{}

func (testType) testMethod() {}

func TestFuncName_Func(t *testing.T) {
	actual := FuncName(testFunc)
	expected := "rft.testFunc"

	if actual != expected {
		t.Errorf(`expected %s but actual is %s`, actual, expected)
	}
}

func TestFuncName_Method(t *testing.T) {
	v := &testType{}

	actual := FuncName(v.testMethod)
	expected := "rft.testType.testMethod-fm"

	if actual != expected {
		t.Errorf(`expected %s but actual is %s`, actual, expected)
	}
}

func TestStructName(t *testing.T) {
	v := &testType{}

	actual := StructName(v)
	expected := "rft.testType"

	if actual != expected {
		t.Errorf(`expected %s but actual is %s`, actual, expected)
	}
}
