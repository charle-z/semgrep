package interfacefuzz
type Bound194 interface{
 ~func() int | ~complex64 | ~struct{ B string } | ~[]byte | Box[string]
}
func Scan194[T Bound194](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
