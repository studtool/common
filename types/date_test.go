package types

import (
	"fmt"
	"testing"
	"time"
)

func TestDate_MarshalJSON(t *testing.T) {
	now := time.Now()
	nowDate := Date(now)

	b, err := nowDate.MarshalJSON()
	if err != nil {
		t.Errorf("Marshalling error: %v", err)
	}

	expected := makeDateString(now)
	actual := string(b)

	if actual != expected {
		t.Errorf("Expected: %s; Actual: %s", expected, actual)
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	now := time.Now()
	nowDate := Date(now)

	expected := makeDateString(now)
	if err := nowDate.UnmarshalJSON([]byte(expected)); err != nil {
		t.Errorf("Unmarshalling error: %v", err)
	}

	actual := makeDateString(time.Time(nowDate))

	if actual != expected {
		t.Errorf("Expected: %s; Actual: %s", actual, expected)
	}
}

func makeDateString(v time.Time) string {
	return fmt.Sprintf(`"%04d-%02d-%02d"`, v.Year(), v.Month(), v.Day())
}
