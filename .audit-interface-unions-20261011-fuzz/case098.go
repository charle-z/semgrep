package interfacefuzz
type Embedded098 interface{~[]rune | ~map[string]rune}
type Bound098 interface{
Embedded098
 ~func(string) bool | ~func() int | ~*string | ~complex64
}
func Scan098[T Bound098](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
