package interfacefuzz
type Bound191 interface{
 ~[3]string | Box[int]
}
func Scan191[T Bound191](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
