package interfacefuzz
type Bound218 interface{
 ~map[string]bool | Pair[string,int] | *int | Box[string]
}
func Scan218[T Bound218](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
