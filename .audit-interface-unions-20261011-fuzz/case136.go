package interfacefuzz
type Bound136 interface{
 ~struct{ A int } | ~[]byte
 M() int
}
func Scan136[T Bound136](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
