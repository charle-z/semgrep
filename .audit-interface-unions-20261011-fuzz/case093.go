package interfacefuzz
import "time"
type Bound093 interface{
 time.Duration | ~map[int]int | string | ~struct{ B string }
}
type Alias093[T Bound093] = []T
func Scan093[T Bound093](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
