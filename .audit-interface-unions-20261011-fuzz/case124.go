package interfacefuzz
type Bound124 interface{
 Box[int] | ~struct{ B string } | **struct{ A int } | Pair[Box[int],Box[string]] | ~map[string]bool
 M() int
}
func Scan124[T Bound124](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
