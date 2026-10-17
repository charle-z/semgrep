package interfacefuzz
type Embedded091 interface{~[]rune | ~map[string]rune}
type Bound091 interface{
Embedded091
 ~chan string | ~func(string) bool
}
func Scan091[T Bound091](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
