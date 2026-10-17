package interfacefuzz
type Bound163 interface{
 **struct{ A int } | ~chan int | ~struct{ A int } | ~map[string]bool
}
func Scan163[T Bound163](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
