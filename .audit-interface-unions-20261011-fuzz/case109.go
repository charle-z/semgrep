package interfacefuzz
type Bound109 interface{
 ~complex64 | ~[]byte | ~func(string) bool | Box[int] | ~[]string
}
func Scan109[T Bound109](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
