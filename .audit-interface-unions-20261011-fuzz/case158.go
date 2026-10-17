package interfacefuzz
type Bound158 interface{
 string | ~[2]int | ~chan int | ~map[string]bool
}
func Scan158[T Bound158](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
