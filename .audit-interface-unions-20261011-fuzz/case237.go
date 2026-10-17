package interfacefuzz
type Bound237 interface{
 ~[]byte | ~[2]int | ~map[string]bool
}
type Alias237[T Bound237] = []T
func Scan237[T Bound237](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
