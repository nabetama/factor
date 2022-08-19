package factor

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFactor(t *testing.T) {
	var tests = []struct {
		given  int
		expect []int
	}{
		{1, []int{1}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{5, []int{5}},
		{7, []int{7}},
		{42, []int{2, 3, 7}},
		{256, []int{2, 2, 2, 2, 2, 2, 2, 2}},
		{1092, []int{2, 2, 3, 7, 13}},
	}

	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.given), func(t *testing.T) {
			if actual := factor(tc.given); !reflect.DeepEqual(actual, tc.expect) {
				t.Errorf("expect %v but %v", tc.expect, actual)
			}
		})
	}
}
