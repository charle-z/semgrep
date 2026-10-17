package interfacefuzz
type Bound113 interface{
 int | ~[]string | ~struct{ A int } | ~complex64
}
func Scan113[T Bound113](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
