package interfacefuzz
type Embedded210 interface{~[]rune | ~map[string]rune}
type Bound210 interface{
Embedded210
 **struct{ A int }
}
type Alias210[T Bound210] = []T
func Scan210[T Bound210](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
