package interfacefuzz
type Bound047 interface{
 ~[]string | Pair[string,int] | **struct{ A int }
}
func Scan047[T Bound047](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
