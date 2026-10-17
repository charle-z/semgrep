package interfacefuzz
type Embedded126 interface{~[]rune | ~map[string]rune}
type Bound126 interface{
Embedded126
 ~map[int]int | Box[int]
}
type Alias126[T Bound126] = []T
func Scan126[T Bound126](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
