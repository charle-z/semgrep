package interfacefuzz
type Bound092 interface{
 Box[string] | ~chan int | ~func() int
 M() int
}
func Scan092[T Bound092](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
