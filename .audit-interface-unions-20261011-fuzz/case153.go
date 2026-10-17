package interfacefuzz
import "time"
type Bound153 interface{
 ~struct{ A int } | ~[]byte | Box[int] | time.Duration
}
type Alias153[T Bound153] = []T
func Scan153[T Bound153](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
