package interfacefuzz
type Bound005 interface{
 ~[3]string
}
func Scan005[T Bound005](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
