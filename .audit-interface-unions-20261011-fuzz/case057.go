package interfacefuzz
type Bound057 interface{
 *int | ~*string | ~map[string]bool
}
type Alias057[T Bound057] = []T
func Scan057[T Bound057](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
