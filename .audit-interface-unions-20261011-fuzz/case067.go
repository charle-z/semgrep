package interfacefuzz
type Bound067 interface{
 ~complex64 | string | ~[]byte
}
func Scan067[T Bound067](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
