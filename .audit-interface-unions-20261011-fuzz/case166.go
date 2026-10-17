package interfacefuzz
import "time"
type Bound166 interface{
 ~[]byte | time.Duration
}
func Scan166[T Bound166](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
