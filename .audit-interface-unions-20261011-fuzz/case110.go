package interfacefuzz
type Bound110 interface{
 Pair[string,int]
}
func Scan110[T Bound110](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
