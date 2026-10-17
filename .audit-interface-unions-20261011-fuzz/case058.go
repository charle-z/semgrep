package interfacefuzz
type Bound058 interface{
 Box[int] | ~struct{ A int } | ~[3]string | ~map[string]bool
}
func Scan058[T Bound058](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
