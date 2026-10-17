package interfacefuzz
type Bound118 interface{
 ~[]string | Pair[string,int] | Box[int] | string
}
func Scan118[T Bound118](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
