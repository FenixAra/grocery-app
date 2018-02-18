package testHelpers

import (
	"reflect"
	"testing"
)

func AssertNil(v interface{}, t *testing.T) {
	if v == nil {
		return
	}
	if reflect.ValueOf(v).Kind().String() == "ptr" && !reflect.ValueOf(v).IsNil() {
		t.Errorf("Expected to be nil | Actual: %+v", v)
		t.FailNow()
	}
}

func AssertNoError(err error, t *testing.T, errMessage string) {
	if err != nil {
		t.Errorf("%s ERR: %v", errMessage, err)
		t.FailNow()
	}
}

func AssertEqual(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
		t.FailNow()
	}
}

func AssertNotEqual(expected interface{}, actual interface{}, t *testing.T) {
	if expected == actual {
		t.Errorf("Assert Failure: Expected Not: %+v | Actual %+v", expected, actual)
		t.FailNow()
	}
}

func AssertNotEmpty(testString string, t *testing.T) {
	if testString == "" {
		t.Errorf("Assert Failure: Expected not to be empty: %+v", testString)
		t.FailNow()
	}
}

func AssertDeepEqual(expected interface{}, actual interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
		t.FailNow()
	}
}

func AssertSliceEqual(expected []interface{}, actual []interface{}, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
		t.FailNow()
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
			t.FailNow()
		}
	}
}
