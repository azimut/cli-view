package tui

import "testing"

func TestRemoveSimilar(t *testing.T) {
	testInputs := []struct {
		input    []string
		expected []string
	}{
		{
			[]string{"http://google.com", "http://google.com/robots.txt"},
			[]string{"http://google.com/robots.txt"},
		},
		{
			[]string{
				"https://www.youtube.com/watch?v=uBsJgceM0KI&ab_channel=Prime",
				"https://www.youtube.com/watch?v=uBsJgceM0KI&ab_channel=PrimeClips",
				"https://www.youtube.com/watch?v=R8-0XZXaJDQ",
			},
			[]string{
				"https://www.youtube.com/watch?v=uBsJgceM0KI&ab_channel=PrimeClips",
				"https://www.youtube.com/watch?v=R8-0XZXaJDQ",
			}},
	}
	var (
		got      []string
		expected []string
	)
	for _, testInput := range testInputs {
		got = removeSimilar(testInput.input)
		expected = testInput.expected
		if !equalSlices(expected, got) {
			t.Errorf("got %v\nexpected %v", got, expected)
		}
	}
}

// equalSlices returns true if both contain the same, in the same order
func equalSlices(as, bs []string) bool {
	if len(as) != len(bs) {
		return false
	}
	for i := range as {
		if as[i] != bs[i] {
			return false
		}
	}
	return true
}
