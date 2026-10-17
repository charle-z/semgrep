package interfacefuzz
type Bound130 interface{
 ~func() int
}
func Scan130[T Bound130](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
