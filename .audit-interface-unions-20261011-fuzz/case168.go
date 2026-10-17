package interfacefuzz
type Embedded168 interface{~[]rune | ~map[string]rune}
type Bound168 interface{
Embedded168
 Pair[int,string] | ~[]string | string | Pair[Box[int],Box[string]]
 M() int
}
type Alias168[T Bound168] = []T
func Scan168[T Bound168](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
