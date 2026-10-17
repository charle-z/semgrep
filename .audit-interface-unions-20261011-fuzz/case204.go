package interfacefuzz
import "time"
type Bound204 interface{
 ~struct{ A int } | ~[]byte | Box[int] | time.Duration
 M() int
}
type Alias204[T Bound204] = []T
func Scan204[T Bound204](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
