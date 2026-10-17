package interfacefuzz
import "time"
type Embedded238 interface{~[]rune | ~map[string]rune}
type Bound238 interface{
Embedded238
 ~struct{ A int } | ~[]byte | Box[int] | time.Duration
}
func Scan238[T Bound238](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
