package interfacefuzz
type Bound211 interface{
 ~struct{ A int } | ~func() int
}
func Scan211[T Bound211](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
