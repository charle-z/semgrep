package interfacefuzz
type Bound229 interface{
 ~complex64 | ~map[string]bool | ~struct{ A int } | Box[string] | *int
}
func Scan229[T Bound229](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
