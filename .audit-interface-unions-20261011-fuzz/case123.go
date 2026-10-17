package interfacefuzz
import "time"
type Bound123 interface{
 ~complex64 | *int | time.Duration | ~func(string) bool
}
type Alias123[T Bound123] = []T
func Scan123[T Bound123](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
