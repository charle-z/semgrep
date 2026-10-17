package interfacefuzz
import "time"
type Bound078 interface{
 **struct{ A int } | ~func(string) bool | time.Duration | ~struct{ B string }
}
type Alias078[T Bound078] = []T
func Scan078[T Bound078](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
