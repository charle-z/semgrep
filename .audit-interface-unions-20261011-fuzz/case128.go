package interfacefuzz
type Bound128 interface{
 ~complex64 | Pair[Box[int],Box[string]] | ~chan string | Box[int]
 M() int
}
func Scan128[T Bound128](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
