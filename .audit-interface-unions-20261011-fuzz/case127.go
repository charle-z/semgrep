package interfacefuzz
type Bound127 interface{
 string | ~*string | Box[string]
}
func Scan127[T Bound127](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
