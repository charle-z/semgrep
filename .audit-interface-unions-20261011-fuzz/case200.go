package interfacefuzz
type Bound200 interface{
 int
 M() int
}
func Scan200[T Bound200](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
