package interfacefuzz
type Bound226 interface{
 ~[2]int | ~chan int
}
func Scan226[T Bound226](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
