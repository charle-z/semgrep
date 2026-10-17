package interfacefuzz
type Bound122 interface{
 string | Box[int] | Pair[string,int]
}
func Scan122[T Bound122](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
