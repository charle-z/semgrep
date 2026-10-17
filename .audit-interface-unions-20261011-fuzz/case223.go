package interfacefuzz
type Bound223 interface{
 ~*string | ~struct{ A int } | ~func() int | ~complex64
}
func Scan223[T Bound223](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
