package interfacefuzz
type Bound233 interface{
 Pair[int,string] | ~[3]string | ~struct{ B string } | ~[]string
}
func Scan233[T Bound233](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
