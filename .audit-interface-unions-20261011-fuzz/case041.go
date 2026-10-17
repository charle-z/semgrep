package interfacefuzz
type Bound041 interface{
 Pair[string,int] | ~struct{ B string }
}
func Scan041[T Bound041](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
