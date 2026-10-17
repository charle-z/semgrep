package interfacefuzz
type Bound096 interface{
 ~struct{ A int } | ~[]string
 M() int
}
type Alias096[T Bound096] = []T
func Scan096[T Bound096](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
