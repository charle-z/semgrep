package interfacefuzz
type Bound152 interface{
 *int | ~*string | ~map[string]bool
 M() int
}
func Scan152[T Bound152](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
