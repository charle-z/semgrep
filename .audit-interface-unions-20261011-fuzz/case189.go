package interfacefuzz
import "time"
type Embedded189 interface{~[]rune | ~map[string]rune}
type Bound189 interface{
Embedded189
 ~complex64 | ~*string | time.Duration | ~[3]string | ~func(string) bool
}
type Alias189[T Bound189] = []T
func Scan189[T Bound189](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
