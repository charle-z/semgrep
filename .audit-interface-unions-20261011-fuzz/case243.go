package interfacefuzz
import "time"
type Bound243 interface{
 ~func(string) bool | *int | ~struct{ B string } | time.Duration
}
type Alias243[T Bound243] = []T
func Scan243[T Bound243](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
