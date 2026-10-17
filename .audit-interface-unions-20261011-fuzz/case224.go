package interfacefuzz
type Embedded224 interface{~[]rune | ~map[string]rune}
type Bound224 interface{
Embedded224
 Pair[Box[int],Box[string]] | ~map[string]bool | *int | string | ~[]string
 M() int
}
func Scan224[T Bound224](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
