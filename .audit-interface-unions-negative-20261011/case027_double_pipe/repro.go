package interfacefuzz
import "time"
type Bound027 interface{
 time.Duration || ~chan string | ~map[string]bool
}
type Alias027[T Bound027] = []T
func Scan027[T Bound027](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
