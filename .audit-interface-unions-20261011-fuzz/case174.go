package interfacefuzz
type Bound174 interface{
 ~map[int]int | int | ~struct{ A int } | ~struct{ B string } | ~[]string
}
type Alias174[T Bound174] = []T
func Scan174[T Bound174](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
