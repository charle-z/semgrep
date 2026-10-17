package interfacefuzz
type Bound169 interface{
 ~chan string | ~func(string) bool | ~struct{ A int } | ~[]byte | **struct{ A int }
}
func Scan169[T Bound169](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
