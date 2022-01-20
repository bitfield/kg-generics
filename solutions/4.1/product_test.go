package product_test

import (
	"product"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestProductInt(t *testing.T) {
	t.Parallel()
	want := 6
	got := product.Product(2, 3)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestProductFloat64(t *testing.T) {
	t.Parallel()
	want := 3.68
	got := product.Product(1.6, 2.3)
	if !cmp.Equal(want, got, cmpopts.EquateApprox(0.0001, 0)) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestProductComplex(t *testing.T) {
	t.Parallel()
	want := 0 + 13i
	got := product.Product(2+3i, 3+2i)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
