package interfacefuzz
type Bound046 interface{
 Pair[Box[int],Box[string]] | ~map[string]bool
}
func Scan046[T Bound046](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
