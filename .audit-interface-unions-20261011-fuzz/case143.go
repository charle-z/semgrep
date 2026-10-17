package interfacefuzz
import "time"
type Bound143 interface{
 ~[]byte /* variant */ | ~map[string]bool /* variant */ | time.Duration /* variant */ | ~func() int
}
func Scan143[T Bound143](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
