package discourse

import (
	"fmt"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
)

func TestFetch(t *testing.T) {
	testUrl := "https://users.rust-lang.org/t/forum-code-formatting-and-syntax-highlighting/42214"
	thread, err := Fetch(testUrl, "Mozilla", time.Second*5)
	if err != nil {
		t.Fail()
	}
	got := thread.id
	expected := 42214
	if got != expected {
		fmt.Println(thread)
		fmt.Println(humanize.Time(thread.op.createdAt))
		t.Errorf("got %d expected %d", got, expected)
	}
}
