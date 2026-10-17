package interfacefuzz
type Bound073 interface{
 ~map[string]bool | ~chan string | ~map[int]int | ~*string
}
func Scan073[T Bound073](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
