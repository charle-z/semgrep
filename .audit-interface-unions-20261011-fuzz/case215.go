package interfacefuzz
type Bound215 interface{
 Pair[Box[int],Box[string]]
}
func Scan215[T Bound215](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
