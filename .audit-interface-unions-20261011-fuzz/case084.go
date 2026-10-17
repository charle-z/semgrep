package interfacefuzz
type Embedded084 interface{~[]rune | ~map[string]rune}
type Bound084 interface{
Embedded084
 ~[2]int | string | ~map[string]bool | ~struct{ A int } | ~[]byte
 M() int
}
type Alias084[T Bound084] = []T
func Scan084[T Bound084](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
