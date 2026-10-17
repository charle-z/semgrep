package interfacefuzz
type Bound144 interface{
 ~func(string) bool | **struct{ A int } | ~chan int | ~map[string]bool | ~map[int]int
 M() int
}
type Alias144[T Bound144] = []T
func Scan144[T Bound144](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
