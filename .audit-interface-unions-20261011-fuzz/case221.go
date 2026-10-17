package interfacefuzz
type Bound221 interface{
 ~struct{ A int } | ~[]byte
}
func Scan221[T Bound221](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
