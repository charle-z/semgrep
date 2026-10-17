package interfacefuzz
type Bound148 interface{
 ~struct{ A int } | ~func() int | ~*string | ~map[string]bool
 M() int
}
func Scan148[T Bound148](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
