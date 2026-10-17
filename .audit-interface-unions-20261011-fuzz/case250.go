package interfacefuzz
type Bound250 interface{
 ~map[string]bool
}
func Scan250[T Bound250](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
