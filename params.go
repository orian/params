package params

import (
	"github.com/julienschmidt/httprouter"

	"strconv"
)

type Params interface {
	Has(name string) bool
	Get(name string) Param
	Add(name string, values ...string)
}

type HttpParams []httprouter.Param

func NewFromHttpRouter(p httprouter.Params) Params {
	return HttpParams(p)
}

func (p HttpParams) Has(name string) bool {
	for _, v := range p {
		if v.Key == name {
			return true
		}
	}
	return false
}

func (p HttpParams) Get(name string) Param {
	var param Param
	for _, v := range p {
		if v.Key == name {
			param = &paramImpl{s: v.Value}
		}
	}
	return param
}

func (p HttpParams) Add(name string, values ...string) {
	panic("cannot modify HttpParams")
}

type paramsImpl map[string]*paramImpl

func NewParams(m map[string]string) Params {
	p := make(paramsImpl)
	for k, v := range m {
		p[k] = &paramImpl{s: v}
	}
	return p
}

func NewParamsSlices(m map[string][]string) Params {
	p := make(paramsImpl)
	for k, v := range m {
		p.Add(k, v...)
	}
	return p
}

func (p paramsImpl) Has(name string) bool {
	_, ok := p[name]
	return ok
}

func (p paramsImpl) Get(name string) Param {
	return p[name]
}

func (p paramsImpl) Add(name string, values ...string) {
	if l := len(values); l == 1 {
		p[name] = &paramImpl{s: values[0]}
	} else if l > 1 {
		p[name] = &paramImpl{ss: append(make([]string, 0, l), values...)}
	}
}

type Param interface {
	CanString() bool
	String() string
	StringOr(v string) string
	toInt(bitSize int)
	CanInt32() bool
	Int32() int32
	Int32Or(v int32) int32
	CanInt64() bool
	Int64() int64
	Int64Or(v int64) int64
	CanInt() bool
	Int() int
	IntOr(v int) int
	toFloat(bitSize int)
	CanFloat32() bool
	Float32() float32
	Float32Or(v float32) float32
	CanFloat64() bool
	Float64() float64
	Float64Or(v float64) float64
	toBool()
	CanBool() bool
	Bool() bool
	BoolOr(v bool) bool
}

type paramImpl struct {
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

func (p *paramImpl) CanString() bool {
	return p != nil
}

func (p *paramImpl) String() string {
	if p.ss == nil {
		return p.s
	}
	return p.ss[0] // TODO what if it's empty?!
}

func (p *paramImpl) StringOr(v string) string {
	if p.CanString() {
		return p.String()
	}
	return v
}

func (p *paramImpl) toInt(bitSize int) {
	p.i, p.e = strconv.ParseInt(p.String(), 10, bitSize)
}

func (p *paramImpl) CanInt32() bool {
	if p == nil {
		return false
	}
	p.toInt(32)
	return p.e == nil
}

func (p *paramImpl) Int32() int32 {
	p.toInt(32)
	return int32(p.i)
}

func (p *paramImpl) Int32Or(v int32) int32 {
	if p.CanInt32() {
		return int32(p.i)
	}
	return v
}

func (p *paramImpl) CanInt64() bool {
	if p == nil {
		return false
	}
	p.toInt(64)
	return p.e == nil
}

func (p *paramImpl) Int64() int64 {
	p.toInt(64)
	return p.i
}

func (p *paramImpl) Int64Or(v int64) int64 {
	if p.CanInt64() {
		return p.i
	}
	return v
}

func (p *paramImpl) CanInt() bool {
	if p == nil {
		return false
	}
	p.toInt(64)
	return p.e == nil
}

func (p *paramImpl) Int() int {
	p.toInt(64)
	return int(p.i)
}

func (p *paramImpl) IntOr(v int) int {
	if p.CanInt() {
		return int(p.i)
	}
	return v
}

func (p *paramImpl) toFloat(bitSize int) {
	p.f, p.e = strconv.ParseFloat(p.String(), bitSize)
}

func (p *paramImpl) CanFloat32() bool {
	if p == nil {
		return false
	}
	p.toFloat(32)
	return p.e == nil
}

func (p *paramImpl) Float32() float32 {
	p.toFloat(32)
	return float32(p.f)
}

func (p *paramImpl) Float32Or(v float32) float32 {
	if p.CanFloat32() {
		return float32(p.f)
	}
	return v
}

func (p *paramImpl) CanFloat64() bool {
	if p == nil {
		return false
	}
	p.toFloat(64)
	return p.e == nil
}

func (p *paramImpl) Float64() float64 {
	p.toFloat(64)
	return p.f
}

func (p *paramImpl) Float64Or(v float64) float64 {
	if p.CanFloat64() {
		return p.f
	}
	return v
}

func (p *paramImpl) toBool() {
	p.b, p.e = strconv.ParseBool(p.String())
}

func (p *paramImpl) CanBool() bool {
	if p == nil {
		return false
	}
	p.toBool()
	return p.e == nil
}

func (p *paramImpl) Bool() bool {
	p.toBool()
	return p.b
}

func (p *paramImpl) BoolOr(v bool) bool {
	if p.CanBool() {
		return p.b
	}
	return v
}
