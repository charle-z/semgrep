package interfacefuzz
type Bound031 interface{
 ~struct{ B string } | Pair[Box[int],Box[string]]
}
func Scan031[T Bound031](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
