package interfacefuzz
type Bound004 interface{
 ~chan string | ~func() int | ~chan int | ~func(string) bool | ~map[string]bool
 M() int
}
func Scan004[T Bound004](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
