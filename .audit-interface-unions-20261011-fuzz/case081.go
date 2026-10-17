package interfacefuzz
type Bound081 interface{
 **struct{ A int } | ~complex64
}
type Alias081[T Bound081] = []T
func Scan081[T Bound081](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
