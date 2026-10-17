package interfacefuzz
type Bound045 interface{
 ~map[string]bool
}
type Alias045[T Bound045] = []T
func Scan045[T Bound045](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
