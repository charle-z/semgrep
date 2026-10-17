package interfacefuzz
type Embedded070 interface{~[]rune | ~map[string]rune}
type Bound070 interface{
Embedded070
 **struct{ A int }
}
func Scan070[T Bound070](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
