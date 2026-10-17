package interfacefuzz
type Bound114 interface{
 *int | ~*string | ~map[string]bool | ~func() int
}
type Alias114[T Bound114] = []T
func Scan114[T Bound114](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
