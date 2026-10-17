package interfacefuzz
type Bound205 interface{
 ~func() int
}
func Scan205[T Bound205](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
