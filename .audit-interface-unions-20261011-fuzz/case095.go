package interfacefuzz
type Bound095 interface{
 *int
}
func Scan095[T Bound095](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
