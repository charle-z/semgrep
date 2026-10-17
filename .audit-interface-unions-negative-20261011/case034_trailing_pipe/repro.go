package interfacefuzz
import "time"
type Bound034 interface{
 ~struct{ A int } | ~[]byte | Box[int] | time.Duration | 
}
func Scan034[T Bound034](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
