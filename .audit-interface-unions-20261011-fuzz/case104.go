package interfacefuzz
type Bound104 interface{
 ~[]byte | ~func(string) bool | Pair[Box[int],Box[string]] | **struct{ A int } | *int
 M() int
}
func Scan104[T Bound104](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
