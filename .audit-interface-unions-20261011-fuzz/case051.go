package interfacefuzz
type Bound051 interface{
 ~struct{ A int } | ~[]byte
}
type Alias051[T Bound051] = []T
func Scan051[T Bound051](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
