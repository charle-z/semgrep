package interfacefuzz
type Bound235 interface{
 ~func() int
}
func Scan235[T Bound235](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
