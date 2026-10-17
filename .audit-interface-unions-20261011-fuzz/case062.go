package interfacefuzz
import "time"
type Bound062 interface{
 ~*string | ~func(string) bool | time.Duration
}
func Scan062[T Bound062](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
