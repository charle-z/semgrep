package interfacefuzz
type Bound178 interface{
 ~chan string | int | ~map[int]int | ~func(string) bool
}
func Scan178[T Bound178](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
