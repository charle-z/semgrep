package interfacefuzz
type Bound225 interface{
 Pair[Box[int],Box[string]]
}
type Alias225[T Bound225] = []T
func Scan225[T Bound225](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
