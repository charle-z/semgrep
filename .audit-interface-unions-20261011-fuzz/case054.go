package interfacefuzz
type Bound054 interface{
 int | Pair[int,string] | ~[]byte | ~*string | Box[int]
}
type Alias054[T Bound054] = []T
func Scan054[T Bound054](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
