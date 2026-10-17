package interfacefuzz
type Bound142 interface{
 ~struct{ A int } | *int | int
}
func Scan142[T Bound142](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
