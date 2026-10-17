package interfacefuzz
import "time"
type Bound179 interface{
 **struct{ A int } | ~map[int]int | ~struct{ A int } | int | time.Duration
}
func Scan179[T Bound179](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
