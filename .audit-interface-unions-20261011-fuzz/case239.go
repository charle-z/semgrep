package interfacefuzz
type Bound239 interface{
 ~complex64 | ~struct{ A int } | ~[]byte | ~func(string) bool | ~chan int
}
func Scan239[T Bound239](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
