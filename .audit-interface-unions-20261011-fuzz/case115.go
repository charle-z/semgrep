package interfacefuzz
type Bound115 interface{
 **struct{ A int }
}
func Scan115[T Bound115](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
