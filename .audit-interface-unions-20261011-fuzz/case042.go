package interfacefuzz
type Embedded042 interface{~[]rune | ~map[string]rune}
type Bound042 interface{
Embedded042
 ~struct{ A int } | Pair[Box[int],Box[string]] | ~map[int]int
}
type Alias042[T Bound042] = []T
func Scan042[T Bound042](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
