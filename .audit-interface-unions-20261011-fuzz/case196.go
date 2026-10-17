package interfacefuzz
import "time"
type Embedded196 interface{~[]rune | ~map[string]rune}
type Bound196 interface{
Embedded196
 time.Duration | ~*string
 M() int
}
func Scan196[T Bound196](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
