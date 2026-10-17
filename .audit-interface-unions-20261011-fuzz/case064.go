package interfacefuzz
type Bound064 interface{
 ~struct{ A int } | ~complex64 | ~*string | string | int
 M() int
}
func Scan064[T Bound064](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
