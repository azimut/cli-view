package discourse

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestToThread(t *testing.T) {
	testUrl := "https://users.rust-lang.org/t/forum-code-formatting-and-syntax-highlighting/42214/2"
	testFile := "../../testdata/discourse-42214.json"
	rawJson, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Error(err)
	}

	thread, err := toThread(string(rawJson), testUrl)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(thread) // DEBUG

	got := thread.id
	expected := 42214
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}

	got = thread.op.id
	expected = 1
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}

	got = len(thread.comments)
	expected = 3
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}
