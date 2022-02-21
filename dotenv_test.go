package dotenv

import (
	"fmt"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	passwordKey := "PASSWORD"
	passwordExpected := "123456"

	Load("./test.env")
	password := os.Getenv(passwordKey)
	if password != passwordExpected {
		t.Error(fmt.Errorf("did not get expected password %v, got %v instead", passwordExpected, password))
	}
}
