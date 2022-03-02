package helpers

import (
	"os"
	"strconv"
	"testing"
)

func TestGetDefaultPort(t *testing.T) {
	defaultPort := 8080
	customPort := 3000
	os.Setenv("SERVER_PORT", "")
	port := GetDefaultPort()
	if port != defaultPort {
		t.Errorf("Expected GetDefaultDatabaseConnectionString() to be equal %d. Got %d", defaultPort, port)
	}

	os.Setenv("SERVER_PORT", strconv.Itoa(customPort))
	port = GetDefaultPort()
	if port != customPort {
		t.Errorf("Expected GetDefaultDatabaseConnectionString() to be equal %d. Got %d", defaultPort, port)
	}
}
