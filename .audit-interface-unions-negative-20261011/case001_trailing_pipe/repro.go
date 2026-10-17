package interfacefuzz
type Bound001 interface{
 ~*string | ~[]string | 
}
func Scan001[T Bound001](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
