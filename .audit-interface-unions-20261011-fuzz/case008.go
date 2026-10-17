package interfacefuzz
type Bound008 interface{
 Pair[int,string] | ~chan int | **struct{ A int } | ~[]string
 M() int
}
func Scan008[T Bound008](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
