package numutil

import "testing"

func TestMax(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, -3}, 2},
	}
	for _, c := range cases {
		got := Max(c.in...)
		if got != c.want {
			t.Errorf("Max(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
