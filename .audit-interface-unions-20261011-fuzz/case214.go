package interfacefuzz
import "time"
type Bound214 interface{
 time.Duration | ~[2]int | ~func(string) bool | Pair[string,int] | ~func() int
}
func Scan214[T Bound214](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
