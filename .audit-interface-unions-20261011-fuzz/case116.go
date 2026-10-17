package interfacefuzz
type Bound116 interface{
 ~[2]int | ~func() int
 M() int
}
func Scan116[T Bound116](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
