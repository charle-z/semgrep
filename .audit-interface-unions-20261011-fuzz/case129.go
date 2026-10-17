package interfacefuzz
type Bound129 interface{
 ~[]byte | Box[int] | ~func() int | ~*string | Pair[Box[int],Box[string]]
}
type Alias129[T Bound129] = []T
func Scan129[T Bound129](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
