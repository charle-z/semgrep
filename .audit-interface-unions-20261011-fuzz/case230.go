package interfacefuzz
type Bound230 interface{
 ~[3]string
}
func Scan230[T Bound230](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
