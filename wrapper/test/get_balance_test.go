package wrapper

import (
	"github.com/refaldyrk/otpweb_wrapper_api/wrapper"
	"testing"
)

func TestGetBalance(t *testing.T) {
	apiKey := ""
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
