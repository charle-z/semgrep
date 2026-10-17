package interfacefuzz
type Bound020 interface{
 ~[]string
 M() int
}
func Scan020[T Bound020](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
