package interfacefuzz
import "time"
type Embedded119 interface{~[]rune | ~map[string]rune}
type Bound119 interface{
Embedded119
 ~struct{ A int } | ~[]byte | Box[int] | time.Duration
}
func Scan119[T Bound119](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
