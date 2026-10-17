package interfacefuzz
type Bound050 interface{
 ~complex64
}
func Scan050[T Bound050](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
