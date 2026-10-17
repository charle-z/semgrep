package interfacefuzz
import "time"
type Bound192 interface{
 Pair[string,int] | ~[3]string | time.Duration
 M() int
}
type Alias192[T Bound192] = []T
func Scan192[T Bound192](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
