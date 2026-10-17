package interfacefuzz
type Bound040 interface{
 ~[3]string
 M() int
}
func Scan040[T Bound040](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
