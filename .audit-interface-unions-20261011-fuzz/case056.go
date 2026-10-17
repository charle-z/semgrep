package interfacefuzz
type Embedded056 interface{~[]rune | ~map[string]rune}
type Bound056 interface{
Embedded056
 ~struct{ A int } | Box[string]
 M() int
}
func Scan056[T Bound056](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
