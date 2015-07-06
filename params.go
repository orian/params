package params

import (
	"strconv"
)

type Params map[string]*Param

func NewParams(m map[string]string) Params {
	p := make(Params)
	for k, v := range m {
		p[k] = &Param{s: v}
	}
	return p
}

func NewParamsSlices(m map[string][]string) Params {
	p := make(Params)
	for k, v := range m {
		if l := len(v); l == 0 {
			continue
		} else if l == 1 {
			p[k] = &Param{s: v[0]}
		} else {
			p[k] = &Param{ss: append(make([]string, 0, l), v...)}
		}
	}
	return p
}

func (p Params) Has(name string) bool {
	_, ok := p[name]
	return ok
}

func (p Params) Get(name string) *Param {
	return p[name]
}

type Param struct {
	s  string
	ss []string
	i  int64
	is []int64
	f  float64
	fs []float64
	b  bool
	bs []bool
	e  error
}

func (p *Param) String() string {
	if p.ss == nil {
		return p.s
	}
	return p.ss[0] // TODO what if it's empty?!
}

func (p *Param) toInt(bitSize int) {
	p.i, p.e = strconv.ParseInt(p.String(), 10, bitSize)
}

func (p *Param) CanInt32() bool {
	if p == nil {
		return false
	}
	p.toInt(32)
	return p.e == nil
}

func (p *Param) Int32() int32 {
	p.toInt(32)
	return int32(p.i)
}

func (p *Param) CanInt64() bool {
	if p == nil {
		return false
	}
	p.toInt(64)
	return p.e == nil
}

func (p *Param) Int64() int64 {
	p.toInt(64)
	return p.i
}

func (p *Param) CanInt() bool {
	if p == nil {
		return false
	}
	p.toInt(64)
	return p.e == nil
}

func (p *Param) Int() int {
	p.toInt(64)
	return int(p.i)
}

func (p *Param) toFloat(bitSize int) {
	p.f, p.e = strconv.ParseFloat(p.String(), bitSize)
}

func (p *Param) CanFloat32() bool {
	if p == nil {
		return false
	}
	p.toFloat(32)
	return p.e == nil
}

func (p *Param) Float32() float32 {
	p.toFloat(32)
	return float32(p.f)
}

func (p *Param) CanFloat64() bool {
	if p == nil {
		return false
	}
	p.toFloat(64)
	return p.e == nil
}

func (p *Param) Float64() float64 {
	p.toFloat(64)
	return p.f
}

func (p *Param) toBool() {
	p.b, p.e = strconv.ParseBool(p.String())
}

func (p *Param) CanBool() bool {
	if p == nil {
		return false
	}
	p.toBool()
	return p.e == nil
}

func (p *Param) Bool() bool {
	p.toBool()
	return p.b
}
