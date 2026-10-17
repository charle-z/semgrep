package interfacefuzz
type Bound030 interface{
 ~[2]int
}
type Alias030[T Bound030] = []T
func Scan030[T Bound030](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
