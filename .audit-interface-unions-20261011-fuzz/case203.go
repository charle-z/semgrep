package interfacefuzz
type Embedded203 interface{~[]rune | ~map[string]rune}
type Bound203 interface{
Embedded203
 ~struct{ B string } | ~map[string]bool | string | ~chan string
}
func Scan203[T Bound203](v T, xs []int){
 auditSink(v)
 for k:=0;k<len(xs);k++ {println(xs[k])}
}
