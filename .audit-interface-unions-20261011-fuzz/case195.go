package interfacefuzz
type Bound195 interface{
 ~map[string]bool
}
type Alias195[T Bound195] = []T
func Scan195[T Bound195](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
