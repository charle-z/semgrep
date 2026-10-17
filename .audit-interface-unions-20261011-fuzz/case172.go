package interfacefuzz
type Bound172 interface{
 ~[]string | ~struct{ A int } | ~func() int
 M() int
}
func Scan172[T Bound172](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
