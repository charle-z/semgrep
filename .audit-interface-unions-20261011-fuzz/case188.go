package interfacefuzz
type Bound188 interface{
 ~complex64 | int | ~func(string) bool | ~[]byte
 M() int
}
func Scan188[T Bound188](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
