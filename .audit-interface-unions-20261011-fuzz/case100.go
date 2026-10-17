package interfacefuzz
type Bound100 interface{
 ~[2]int
 M() int
}
func Scan100[T Bound100](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
