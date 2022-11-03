package pivot_server

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("assertion failed: want %v got %v", expected, actual)
	}
}
func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("assertion failed: want no error, got %v", err)
	}
}

func assertStatusCode(t *testing.T, rr *httptest.ResponseRecorder, expected int) {
	if rr.Result().StatusCode != expected {
		t.Errorf("wrong status code: want %d got %d", expected, rr.Result().StatusCode)
	}
}
