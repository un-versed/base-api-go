package helpers

import (
	"os"
	"strconv"
	"testing"
)

func TestGetDefaultDatabaseConnectionString(t *testing.T) {
	defaultString := "postgres://postgres:postgres@localhost:5432/postgres"
	customString := "postgres://test:test@localhost:5432/postgres"
	os.Setenv("DATABASE_URL", "")
	connString := GetDefaultDatabaseConnectionString()
	if connString != defaultString {
		t.Errorf("Expected GetDefaultDatabaseConnectionString() to be equal %s. Got %s", defaultString, connString)
	}

	os.Setenv("DATABASE_URL", customString)
	connString = GetDefaultDatabaseConnectionString()
	if connString != customString {
		t.Errorf("Expected GetDefaultDatabaseConnectionString() to be equal %s. Got %s", defaultString, connString)
	}
}

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
