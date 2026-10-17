package interfacefuzz
type Embedded140 interface{~[]rune | ~map[string]rune}
type Bound140 interface{
Embedded140
 ~map[string]bool
 M() int
}
func Scan140[T Bound140](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
