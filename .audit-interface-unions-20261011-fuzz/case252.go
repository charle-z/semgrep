package interfacefuzz
type Embedded252 interface{~[]rune | ~map[string]rune}
type Bound252 interface{
Embedded252
 ~struct{ B string } | Pair[Box[int],Box[string]] | ~[3]string
 M() int
}
type Alias252[T Bound252] = []T
func Scan252[T Bound252](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
