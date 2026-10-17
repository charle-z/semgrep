package interfacefuzz
type Bound090 interface{
 Box[int]
}
type Alias090[T Bound090] = []T
func Scan090[T Bound090](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
