package buildapileofcubes

import "testing"

func TestFindNb(t *testing.T) {
	t.Run("The cube exists", func(t *testing.T) {
		got := FindNb(1071225)
		want := 45

		if got != want {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("The cube does not exist", func(t *testing.T) {
		got := FindNb(91716553919377)
		want := -1

		if got != want {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})
}
