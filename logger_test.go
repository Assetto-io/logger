package logger

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	l, err := New("test_service", "info")
	if err != nil {
		t.Errorf("Logger is %v; wanted: smth else", l)
	}
}

func TestField(t *testing.T) {
	testValues := []string{"a", "b", "c", "d"}
	got := Field("some values", testValues)
	if got.Key != "some values" {
		t.Errorf("Key is %v; wanted: some values", got)
	}
	if err := testField(got.Value, testValues); err != nil {
		t.Errorf(err.Error())
	}
}

func TestLogger_Info(t *testing.T) {
	l, err := New("test_service", "debug")
	if err != nil {
		t.Errorf("Logger is %v; wanted: smth else", l)
	}
	l.Info("some logs", Field("data", 5))
	l.Debug("some logs", Field("data", 5))
	l.Error("some logs", errors.New("Some custom error"), Field("data", 5))
}

// ======================== Helper Functions ========================

func testField(t interface{}, v []string) error {
	switch t := t.(type) {
	case []string:
		for i, value := range t {
			if value != v[i] {
				return errors.New("Wrong value")
			}
		}
	default:
		return errors.New("Wrong value")
	}
	return nil
}
