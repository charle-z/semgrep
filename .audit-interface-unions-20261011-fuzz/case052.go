package interfacefuzz
type Bound052 interface{
 ~complex64 | ~struct{ A int } | ~func() int
 M() int
}
func Scan052[T Bound052](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
