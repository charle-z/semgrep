package interfacefuzz
type Bound106 interface{
 ~chan int | Pair[Box[int],Box[string]]
}
func Scan106[T Bound106](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
