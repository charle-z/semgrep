package interfacefuzz
import "time"
type Bound101 interface{
 time.Duration | ~[]byte
}
func Scan101[T Bound101](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
