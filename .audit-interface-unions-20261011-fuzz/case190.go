package interfacefuzz
type Bound190 interface{
 *int
}
func Scan190[T Bound190](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
