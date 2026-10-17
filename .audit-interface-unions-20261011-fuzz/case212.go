package interfacefuzz
type Bound212 interface{
 ~func() int | ~[]string | Pair[int,string]
 M() int
}
func Scan212[T Bound212](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
