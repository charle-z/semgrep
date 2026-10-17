package interfacefuzz
type Bound072 interface{
 ~struct{ A int } | ~func() int | Box[int]
 M() int
}
type Alias072[T Bound072] = []T
func Scan072[T Bound072](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
