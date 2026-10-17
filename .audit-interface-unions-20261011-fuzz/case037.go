package interfacefuzz
type Bound037 interface{
 *int | Pair[Box[int],Box[string]] | Box[int]
}
func Scan037[T Bound037](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
