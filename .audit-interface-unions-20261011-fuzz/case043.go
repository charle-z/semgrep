package interfacefuzz
type Bound043 interface{
 Box[string] | ~complex64 | ~map[int]int | ~map[string]bool
}
func Scan043[T Bound043](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
