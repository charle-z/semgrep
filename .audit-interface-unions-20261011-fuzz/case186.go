package interfacefuzz
type Bound186 interface{
 Pair[string,int] | int
}
type Alias186[T Bound186] = []T
func Scan186[T Bound186](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
