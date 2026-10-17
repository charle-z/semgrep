package interfacefuzz
type Bound208 interface{
 ~[]byte | Box[int] | ~struct{ A int } | ~struct{ B string }
 M() int
}
func Scan208[T Bound208](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
