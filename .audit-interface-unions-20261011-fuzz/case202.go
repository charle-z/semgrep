package interfacefuzz
type Bound202 interface{
 ~[2]int | ~chan int | *int
}
func Scan202[T Bound202](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
