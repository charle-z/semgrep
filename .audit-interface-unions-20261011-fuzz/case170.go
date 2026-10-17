package interfacefuzz
type Bound170 interface{
 ~struct{ A int }
}
func Scan170[T Bound170](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
