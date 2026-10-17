package interfacefuzz
type Bound247 interface{
 *int | ~*string | ~map[string]bool
}
func Scan247[T Bound247](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
