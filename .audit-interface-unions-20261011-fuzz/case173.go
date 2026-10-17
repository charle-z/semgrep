package interfacefuzz
type Bound173 interface{
 Pair[string,int] | **struct{ A int } | ~complex64 | ~[]byte
}
func Scan173[T Bound173](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
