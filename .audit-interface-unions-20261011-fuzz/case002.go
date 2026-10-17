package interfacefuzz
type Bound002 interface{
 int | **struct{ A int } | ~[3]string
}
func Scan002[T Bound002](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
