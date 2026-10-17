package interfacefuzz
import "time"
type Bound089 interface{
 string | ~chan int | time.Duration | ~chan string | Pair[int,string]
}
func Scan089[T Bound089](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
