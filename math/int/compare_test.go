package int

import "testing"

func TestMax(t *testing.T) {
	if Max(2,3) != 3 {
		t.Fatalf("failed get max num")
	}
}

func TestMin(t *testing.T) {
	if Min(2,3) != 2 {
		t.Fatalf("failed get min num")
	}
}
