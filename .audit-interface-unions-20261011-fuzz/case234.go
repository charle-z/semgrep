package interfacefuzz
type Bound234 interface{
 ~struct{ A int } | ~chan string | ~map[string]bool | Pair[string,int] | ~func() int
}
type Alias234[T Bound234] = []T
func Scan234[T Bound234](v T, xs []int){
 auditSink(v)
 for range xs {println("seen")}
}
