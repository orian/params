package params

import (
	"github.com/julienschmidt/httprouter"

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
	if m := p.Get("random"); m.CanBool() {
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

func TestParseOrDef(t *testing.T) {
	x := map[string]string{"city": "Olsztyn", "year": "2015", "badF": "x"}
	p := NewParams(x)

	if x := p.Get("year").IntOr(1234); x != 2015 {
		t.Errorf("want: 2015, got: %d", x)
	}
	if x := p.Get("badF").Float32Or(3.14); x != 3.14 {
		t.Errorf("want: 3.14, got: %f", x)
	}
	if x := p.Get("empty").Float32Or(3.14); x != 3.14 {
		t.Errorf("want: 3.14, got: %f", x)
	}
}

func TestStringFromHttprouter(t *testing.T) {
	p := make(httprouter.Params, 0)
	p = append(p, httprouter.Param{"name", "value"})

	m := NewFromHttpRouter(p)
	s := m.Get("name").String()
	if s != "value" {
		t.Errorf("want: `value`, got: %q", s)
	}
}

func TestInt64FromHttprouter(t *testing.T) {
	p := make(httprouter.Params, 0)
	p = append(p, httprouter.Param{"name", "3"})

	m := NewFromHttpRouter(p)
	s := m.Get("name").Int64Or(1)
	if s != 3 {
		t.Errorf("want: 3, got: %d", s)
	}
}
