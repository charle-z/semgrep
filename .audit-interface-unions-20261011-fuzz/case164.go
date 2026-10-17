package interfacefuzz
type Bound164 interface{
 ~struct{ B string } | **struct{ A int } | ~struct{ A int } | ~complex64 | ~func() int
 M() int
}
func Scan164[T Bound164](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
