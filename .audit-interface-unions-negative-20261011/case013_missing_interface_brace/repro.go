package interfacefuzz
type Bound013 interface[
 ~[]byte | ~struct{ A int } | ~func(string) bool | string
}
func Scan013[T Bound013](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
