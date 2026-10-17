package interfacefuzz
type Bound085 interface{
 ~struct{ A int }
}
func Scan085[T Bound085](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
