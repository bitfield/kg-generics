package point_test

import (
	"point"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetX_ReturnsXValueOfPoint(t *testing.T) {
	t.Parallel()
	p := point.Point{
		X: 1,
		Y: 2,
	}
	want := 1
	got := point.GetX[point.Point](p)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetY_ReturnsYValueOfPoint(t *testing.T) {
	t.Parallel()
	p := point.Point{
		X: 1,
		Y: 2,
	}
	want := 2
	got := point.GetY[point.Point](p)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
