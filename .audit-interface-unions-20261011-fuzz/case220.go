package interfacefuzz
type Bound220 interface{
 ~*string
 M() int
}
func Scan220[T Bound220](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
