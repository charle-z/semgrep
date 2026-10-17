package interfacefuzz
type Bound102 interface{
 ~struct{ A int } | ~[]byte | Box[int]
}
type Alias102[T Bound102] = []T
func Scan102[T Bound102](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
