package permutationinstring

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "Example 1",
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			name: "Example 2",
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
