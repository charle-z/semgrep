package interfacefuzz
type Bound146 interface{
 ~func() int | ~struct{ A int }
}
func Scan146[T Bound146](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
