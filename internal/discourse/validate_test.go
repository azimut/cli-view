package discourse

import "testing"

func TestEffectiveUrl(t *testing.T) {
	testUrls := map[string]string{
		"https://llllllll.co/t/disquiet-junto-project-0590-concrete-roots/62027":       "https://llllllll.co/t/disquiet-junto-project-0590-concrete-roots/62027.json",
		"https://0x00sec.org/t/question-on-your-need-to-be-known/34563":                "https://0x00sec.org/t/question-on-your-need-to-be-known/34563.json",
		"https://discourse.haskell.org/t/an-opportunity-that-i-couldnt-pass-up/7485/":  "https://discourse.haskell.org/t/an-opportunity-that-i-couldnt-pass-up/7485.json",
		"https://discourse.haskell.org/t/an-opportunity-that-i-couldnt-pass-up/7485/3": "https://discourse.haskell.org/t/an-opportunity-that-i-couldnt-pass-up/7485.json",
	}
	for rawUrl, expected := range testUrls {
		got, _ := effectiveUrl(rawUrl)
		if expected != got {
			t.Errorf("got %s expected %s", got, expected)
		}
	}
}
