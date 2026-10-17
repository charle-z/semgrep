package interfacefuzz
type Bound137 interface{
 ~[3]string | ~[2]int | Pair[Box[int],Box[string]]
}
func Scan137[T Bound137](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
