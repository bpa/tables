package auth

import (
	"testing"

	"github.com/bpa/tables/test"
)

func TestAuth(t *testing.T) {
	c := test.NewClient(test.Ignore)
	_, err := Login(c, map[string]interface{}{"cmd": "login"})
	if err == nil {
		t.Error("Expected error, got none")
	} else {
		if err.Error() != "Login missing field: `method`" {
			t.Errorf("Expected error `Login missing type`, got `%s`", err.Error())
		}
	}
}
