package interfacefuzz
import "time"
type Bound025 interface{
 time.Duration
}
func Scan025[T Bound025](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
