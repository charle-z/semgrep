package interfacefuzz
type Bound222 interface{
 ~func() int | Box[string] | ~[2]int
}
type Alias222[T Bound222] = []T
func Scan222[T Bound222](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
