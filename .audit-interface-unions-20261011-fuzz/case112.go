package interfacefuzz
type Embedded112 interface{~[]rune | ~map[string]rune}
type Bound112 interface{
Embedded112
 ~complex64 | ~func(string) bool | ~map[int]int
 M() int
}
func Scan112[T Bound112](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
