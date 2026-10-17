package interfacefuzz
type Bound015 interface{
 ~map[string]bool
}
type Alias015[T Bound015] = []T
func Scan015[T Bound015](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
