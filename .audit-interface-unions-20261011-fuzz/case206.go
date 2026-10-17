package interfacefuzz
type Bound206 interface{
 int | *int
}
func Scan206[T Bound206](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
