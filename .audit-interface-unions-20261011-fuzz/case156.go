package interfacefuzz
type Bound156 interface{
 ~chan string | ~chan int
 M() int
}
type Alias156[T Bound156] = []T
func Scan156[T Bound156](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
