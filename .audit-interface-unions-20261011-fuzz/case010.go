package interfacefuzz
type Bound010 interface{
 Pair[string,int]
}
func Scan010[T Bound010](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
