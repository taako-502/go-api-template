package page

import "testing"

func TestMap(t *testing.T) {
	got := (Page[int]{Items: []int{1, 2}, Total: 2}).Map(func(v int) int { return v * 2 })
	if got.Total != 2 || got.Items[0] != 2 || got.Items[1] != 4 {
		t.Fatalf("got=%+v", got)
	}
}
