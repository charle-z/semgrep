package interfacefuzz
type Bound065 interface{
 Pair[Box[int],Box[string]]
}
func Scan065[T Bound065](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
