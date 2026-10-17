package interfacefuzz
import "time"
type Bound024 interface{
 ~func(string) bool | time.Duration | ~struct{ A int } | string | Box[int]
 M() int
}
type Alias024[T Bound024] = []T
func Scan024[T Bound024](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
