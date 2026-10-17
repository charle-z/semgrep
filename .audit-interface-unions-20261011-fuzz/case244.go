package interfacefuzz
type Bound244 interface{
 ~func(string) bool | ~func() int | ~struct{ A int } | ~chan string | Pair[int,string]
 M() int
}
func Scan244[T Bound244](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
