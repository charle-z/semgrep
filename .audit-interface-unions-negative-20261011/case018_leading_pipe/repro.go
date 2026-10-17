package interfacefuzz
import "time"
type Bound018 interface{
 |  time.Duration | ~struct{ A int } | **struct{ A int } | Pair[string,int]
}
type Alias018[T Bound018] = []T
func Scan018[T Bound018](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
