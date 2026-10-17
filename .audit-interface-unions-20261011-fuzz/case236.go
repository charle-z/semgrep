package interfacefuzz
type Bound236 interface{
 **struct{ A int } | Pair[string,int]
 M() int
}
func Scan236[T Bound236](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
