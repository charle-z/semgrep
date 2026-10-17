package interfacefuzz
type Embedded014 interface{~[]rune | ~map[string]rune}
type Bound014 interface{
Embedded014
 ~complex64 | ~[2]int | **struct{ A int } | Pair[Box[int],Box[string]] | *int
}
func Scan014[T Bound014](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
