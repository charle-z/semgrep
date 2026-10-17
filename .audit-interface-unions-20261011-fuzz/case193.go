package interfacefuzz
import "time"
type Bound193 interface{
 ~[]string | time.Duration | ~[3]string | ~func(string) bool
}
func Scan193[T Bound193](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
