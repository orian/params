# params
A simple wrapper to parse `map[string]string` and `map[string][]string` values as int, float, bool.

It was written with Gorilla Mux (https://github.com/gorilla/mux) and `url.Values` in mind.

#API

There are two types `Params` and `Param`. The first one is a collection of all params. 
To create a `Params` one should call one of:
 
  - `func NewParams(m map[string]string) Params`
  - `func NewParamsSlices(m map[string][]string) Params`

One can check if a param is available and get it using: `func (*Params) Has(name string) bool` and `func (*Params) Get(name string) *Param`. `Param` keeps value of one param. Then one may check if parsing param as a given type is possible using one of many functions: `func (*Param) CanInt() bool` (similar for `int32`, `int64`, etc.).
One may get value as a given type by: `func (*Param) Float32() float32`, it will return a value parsed as given type. It will return default value if cannot parse. 

**Beware**, the code:
```go
ps := params.NewParams(...).Get("var")
p := ps.Float32()
```
will panic if Params doesn't have a `var` parameter set. It's because `Get` will return nil and one calls a method on `nil`. It's by design. If one wants default's then should use `.Float32Or(0.)` which if parsing not possible or with error returns an argument.

# Example

```go
import (
  "github.com/orian/params"
  "github.com/gorilla/mux"
)
```
And
```go
func ListH(w http.ResponseWriter, req *http.Request) {
	page := params.NewParams(mux.Vars(req)).Get("page").IntOr(0)
	fmt.Fprintf(w, "List view, page: %d", page)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/new", FormH).Methods("GET")
	r.HandleFunc("/upload", UploadH).Methods("POST")
	r.HandleFunc("/", ListH)
	r.HandleFunc("/p/{page}", ListH)
	r.HandleFunc("/k/{id}", DetailH)

	log.Print(http.ListenAndServe(":8080", r))
}

```
