package interfacefuzz
import "time"
type Bound254 interface{
 ~func(string) bool | time.Duration | ~complex64 | Pair[string,int] | ~chan int
}
func Scan254[T Bound254](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
