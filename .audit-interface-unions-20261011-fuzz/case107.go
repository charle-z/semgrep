package interfacefuzz
import "time"
type Bound107 interface{
 **struct{ A int } | time.Duration | Box[int]
}
func Scan107[T Bound107](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
