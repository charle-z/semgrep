package interfacefuzz
type Bound080 interface{
 ~complex64
 M() int
}
func Scan080[T Bound080](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
