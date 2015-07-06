package params

import (
	"testing"
)

func TestSimple(t *testing.T) {
	x := map[string]string{"city": "Olsztyn"}
	p := NewParams(x)
	k := "city"
	if !p.Has(k) {
		t.Errorf("expecting to have %s key", k)
	}
	if m := p.Get(k); m.CanBool() || m.CanFloat32() || m.CanFloat64() ||
		m.CanInt() || m.CanInt32() || m.CanInt64() {
		t.Error("not exppected conversion")
	} else if m.String() != "Olsztyn" {
		t.Errorf("want: Olsztyn, got: %s", m.String())
	}
}

func TestEmpty(t *testing.T) {
	x := map[string]string{"city": "Olsztyn"}
	p := NewParams(x)
	if m := p.Get("random"); m != nil {
		t.Error("nil expected")
	} else if m.CanBool() {
		t.Error("empty should not be convertable")
	}
}

func TestFromMapOfSlice(t *testing.T) {
	x := map[string][]string{"city": []string{"Olsztyn"}}
	p := NewParamsSlices(x)

	if m := p.Get("city"); m == nil {
		t.Error("expected not nil")
	} else if m.String() != "Olsztyn" {
		t.Errorf("want: Olszty, got: %s", m.String())
	}
}
