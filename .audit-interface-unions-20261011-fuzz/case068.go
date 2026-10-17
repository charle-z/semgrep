package interfacefuzz
import "time"
type Bound068 interface{
 ~struct{ A int } | ~[]byte | Box[int] | time.Duration
 M() int
}
func Scan068[T Bound068](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
