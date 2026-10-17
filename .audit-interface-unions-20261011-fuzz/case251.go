package interfacefuzz
type Bound251 interface{
 string | Box[string]
}
func Scan251[T Bound251](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
