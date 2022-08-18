package factor

import "testing"

func TestFactor(t *testing.T) {
	var tests = []struct{
		name string
		given int
		expect string
	}{
		{"12 -> 2 2 3", 12, "2 2 3"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if actual := factor(tc.given); actual != tc.expect {
				t.Errorf("expect %s but %s", tc.expect, actual)
			}
		})
	}
}