package interfacefuzz
type Bound125 interface{
 **struct{ A int }
}
func Scan125[T Bound125](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
