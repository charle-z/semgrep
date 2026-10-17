package interfacefuzz
type Bound038 interface{
 *int | ~*string | ~map[string]bool | ~func() int
}
func Scan038[T Bound038](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
