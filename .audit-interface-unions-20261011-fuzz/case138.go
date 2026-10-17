package interfacefuzz
type Bound138 interface{
 ~[2]int | Box[string] | ~*string | int
}
type Alias138[T Bound138] = []T
func Scan138[T Bound138](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
