package lobsters

import (
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	// url := "https://lobste.rs/s/jclvos/one_decade_lobsters"
	url := "https://lobste.rs/s/wucxh0/mitchell_hashimoto_uses_simple_code"
	_, err := Fetch(url, "LobsterView/1.0", 5*time.Second)
	if err != nil {
		t.Errorf("failed to retrieve url")
	}
}
