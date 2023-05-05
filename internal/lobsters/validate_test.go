package lobsters

import "testing"

func TestEffectiveUrlPositive(t *testing.T) {
	urls := map[string]string{
		"https://lobste.rs/s/ab6gsq/maybe_you_should_store_passwords": "https://lobste.rs/s/ab6gsq.json",
		"https://lobste.rs/s/ab6gsq/":                                 "https://lobste.rs/s/ab6gsq.json",
		"https://lobste.rs/s/ab6gsq":                                  "https://lobste.rs/s/ab6gsq.json",
	}
	for rawUrl, expected := range urls {
		got, err := effectiveUrl(rawUrl)
		if err != nil {
			t.Fatalf("could not parse url %s with error: %s", rawUrl, err)
		}
		if expected != got {
			t.Errorf("got %s expected %s", got, expected)
		}
	}
}

func TestEffectiveUrlNegative(t *testing.T) {
	urls := map[string]error{
		"https://www.qword.net/2023/04/30/maybe-you-should-store-passwords-in-plaintext": ErrInvalidDomain,
	}
	for rawUrl, expectedErr := range urls {
		_, err := effectiveUrl(rawUrl)
		if err != expectedErr {
			t.Errorf("did not errored for %s", rawUrl)
		}
	}
}
