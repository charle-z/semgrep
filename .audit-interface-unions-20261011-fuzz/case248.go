package interfacefuzz
type Bound248 interface{
 Box[int] | ~struct{ A int } | **struct{ A int } | Pair[string,int]
 M() int
}
func Scan248[T Bound248](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
