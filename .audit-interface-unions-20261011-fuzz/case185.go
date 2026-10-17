package interfacefuzz
type Bound185 interface{
 ~map[string]bool
}
func Scan185[T Bound185](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
