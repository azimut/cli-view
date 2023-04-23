package vichan

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	rawUrls := []string{
		"https://lainchan.org/%CE%BB/res/35065.html",
		"https://wired-7.org/tech/res/5297.html",
	}
	for _, rawUrl := range rawUrls {
		effectiveUrl, err := parseUrl(rawUrl)
		if err != nil {
			t.Errorf("parsing %s error'ed with %v", rawUrl, err)
		}
		if !strings.HasSuffix(effectiveUrl, ".json") {
			t.Errorf("effectiveUrl missing .json suffix %s", effectiveUrl)
		}
	}
	rawUrls = []string{
		"",
		"https://wired-7.org/tech/res/",
	}
	for _, rawUrl := range rawUrls {
		_, err := parseUrl(rawUrl)
		if err == nil {
			t.Errorf("this should have failed, but it didn't %s", rawUrl)
		}
	}
}
