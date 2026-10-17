package interfacefuzz
type Bound167 interface{
 Box[string] | ~[]string | Pair[string,int]
}
func Scan167[T Bound167](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
