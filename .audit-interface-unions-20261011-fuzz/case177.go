package interfacefuzz
import "time"
type Bound177 interface{
 string | ~struct{ A int } | time.Duration
}
type Alias177[T Bound177] = []T
func Scan177[T Bound177](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
