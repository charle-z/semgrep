package interfacefuzz
import "time"
type Bound083 interface{
 time.Duration | **struct{ A int } | ~[3]string | ~[]byte
}
func Scan083[T Bound083](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
