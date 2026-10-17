package interfacefuzz
import "time"
type Bound151 interface{
 ~complex64 | time.Duration
}
func Scan151[T Bound151](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
