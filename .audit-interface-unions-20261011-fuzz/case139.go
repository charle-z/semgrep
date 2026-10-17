package interfacefuzz
type Bound139 interface{
 int | *int | ~map[int]int | ~[]byte | ~func(string) bool
}
func Scan139[T Bound139](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
