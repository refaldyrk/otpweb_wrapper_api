package wrapper

import (
	"github.com/refaldyrk/otpweb_wrapper_api/wrapper"
	"testing"
)

func TestGetBalance(t *testing.T) {
	apiKey := "d9f09e339e85578a1932dd09fdec15435b486c30"
	w := wrapper.NewWrapper(apiKey)

	result, err := w.GetBalance()

	if err != nil {
		t.Errorf("GetBalance returned an error: %v", err)
	}

	if !result.Status {
		t.Errorf("Got False Status")
	}

	t.Logf("Balance: %+v", result)
}
