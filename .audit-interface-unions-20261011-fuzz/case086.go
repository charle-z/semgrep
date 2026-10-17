package interfacefuzz
type Bound086 interface{
 ~chan string | Pair[int,string]
}
func Scan086[T Bound086](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
