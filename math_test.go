package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
