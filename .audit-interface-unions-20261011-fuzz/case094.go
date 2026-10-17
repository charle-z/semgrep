package interfacefuzz
type Bound094 interface{
 ~*string | ~[3]string | Pair[Box[int],Box[string]] | ~chan string | ~chan int
}
func Scan094[T Bound094](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
