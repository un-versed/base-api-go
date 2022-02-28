package models

import (
	"reflect"
	"testing"
)

func TestErrorModel_NewApiError(t *testing.T) {
	t.Run("API Error 1", auxTestErrorModel_NewApiError("[API] Test Error", &ApiError{Message: "[API] Test Error"}))
}

func auxTestErrorModel_NewApiError(m string, expected *ApiError) func(*testing.T) {
	return func(t *testing.T) {
		errorMessage := NewApiError(m)

		if !reflect.DeepEqual(errorMessage, expected) {
			t.Errorf("Expected NewApiError(m: %s) to be equal %+v. Got %+v", m, expected, errorMessage)
		}
	}
}
